# Ревью последних изменений

**Коммит:** `6d508eb` — *before pizdul ot vladosa*
**Ветка:** `develop`
**Дата ревью:** 2026-04-21
**Итог:** проект собирается (`go build ./...` проходит, `go vet` чист), но в коде боевые баги и серьёзный архитектурный регресс.

---

## 🔴 Критические баги

### 1. `rows.Close()` не вызывается при успехе — утечка соединений пула

**Файлы:**
- `internal/repository/storage/login/methods.go:20-35`
- `internal/repository/storage/test/methods.go:9-40`

```go
rows, err := l.db.Query(query, ...)
if err != nil {
    defer rows.Close()   // defer внутри if — бессмысленно, rows == nil
}
for rows.Next() { ... }  // при err != nil будет nil pointer panic
// rows.Close() на успешном пути НЕ вызывается
```

**Последствия:**
- При ошибке запроса — паника на `rows.Next()` (обращение к nil).
- При успехе — `rows.Close()` не вызывается, соединение не возвращается в пул. Учитывая `db.SetMaxOpenConns(25)` в `internal/postgres/db.go:29` — через 25 запросов сервер заблокируется.

**Как должно быть:**
```go
rows, err := l.db.QueryContext(ctx, query, ...)
if err != nil {
    return loginResponse, err
}
defer rows.Close()
```

---

### 2. `Password int` — провал по безопасности и функциональности

**Файлы:**
- `internal/model/model.go:8`
- `internal/model/modelLogin.go:11`

```go
type User struct {
    ID       int    `json:"user_id"`
    Login    string `json:"loginZ"`    // опечатка?
    Password int    `json:"passwordZ"` // ← проблема
}
type LoginRequest struct {
    Login    string `json:"login"`
    Password int    `json:"password"`  // ← проблема
}
```

**Проблемы:**
- Пароль с любой буквой/символом не принимается — `ShouldBindJSON` упадёт на `"password": "qwerty"`.
- В SQL (`login/methods.go:17`) сравнение `WHERE u.password = $2` — plain-текст без хеша.
- JSON-теги `loginZ`, `passwordZ` — похоже на опечатки/заглушки, API с такими именами выглядит сломанным.

**Как должно быть:** `Password string` + bcrypt-хеш при сохранении и `bcrypt.CompareHashAndPassword` при проверке.

---

### 3. Хардкод Windows-пути в конфиге

**Файл:** `internal/config/config.go:41`

```go
pathEnv := path.Join("D:/ApplicationBackend/.env")
err := godotenv.Load(pathEnv)
if err != nil {
    return nil, fmt.Errorf("No .env file found, using system environment variables")
}
```

**Проблемы:**
- На macOS/Linux всегда вернёт ошибку, `Load()` упадёт.
- `path.Join` с одним аргументом — бессмысленный вызов.
- Текст ошибки говорит «using system environment variables», но на самом деле это **fatal error** — fallback не реализован.

**Аналогично** в `.vscode/launch.json:9` — `"program": "D:/ApplicationBackend/cmd/main.go"`.

**Как должно быть:** читать путь из переменной окружения или относительно `os.Getwd()`, при отсутствии `.env` — честный fallback на `os.Getenv`.

---

### 4. `loginService.SignIn` глотает реальную ошибку БД

**Файл:** `internal/service/loginservice/service_login.go:11-17`

```go
response, err := s.repository.SignIn(ctx, loginRequest)

if response.Id == 0 && response.Name == "" && response.Role == "" {
    return response, errors.New("User not found")  // err от репо затёрт
}
return response, err
```

**Проблемы:**
- Если БД упала или сеть легла — пользователь увидит «User not found» вместо 500.
- В `internal/model/errors.go` объявлен `ErrUserNotFound`, но **не используется**. Создаётся новая ошибка через `errors.New`.

**Как должно быть:**
```go
response, err := s.repository.SignIn(ctx, loginRequest)
if err != nil {
    return response, err
}
if response.Id == 0 {
    return response, model.ErrUserNotFound
}
return response, nil
```

---

### 5. Проигнорированная ошибка + N+1 в `GetAvailableTests`

**Файл:** `internal/service/serviceT/service.go:12-27`

```go
tests, err := t.test.GetAvailableTestsId(ctx, user_id)
// err здесь НЕ проверяется!
var avaiilableTests []model.TestFull
for _, value := range tests {
    test, err := t.GetTestById(ctx, value)
    if err != nil { return nil, err }
    if test.Id != 0 && test.Description != "" && test.Title != "" {
        avaiilableTests = append(avaiilableTests, test)
    }
}
return avaiilableTests, err
```

**Проблемы:**
- **Игнор ошибки** от `GetAvailableTestsId` — цикл пойдёт по nil/пустому срезу, а в конце вернётся именно эта потерянная `err`.
- **N+1**: для каждого теста отдельный запрос в БД. В удалённом `testRepositoryImt.go` это было решено одним JOIN.
- Опечатка `avaiilableTests` → `availableTests`.
- Проверка «непустоты» полей — костыль; лучше на уровне SQL гарантировать валидность.

---

## 🟠 Архитектурный регресс

### Удалены все интерфейсы

**Удалены файлы:**
- `internal/repository/repository.go` (интерфейсы `UserRepository`, `TestRepository`, `LoginRepository`)
- `internal/repository/storage/interface.go`

**Было:**
```go
type TestRepository interface {
    GetTestById(ctx context.Context, id int) (model.TestFull, error)
    GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull, error)
}
func NewTestService(testRepo repository.TestRepository) TestService
```

**Стало:**
```go
type TestService struct {
    test *test.TestStorage  // конкретный тип из пакета storage
}
```

**Последствия:**
- **Невозможно писать юнит-тесты на сервисный слой** — без мока репозитория нужен поднятый PostgreSQL для каждого теста.
- Нарушен Dependency Inversion Principle — высокоуровневая логика (service) зависит от низкоуровневой (storage).
- В `cmd/main.go:32` сам разработчик оставил маркер: `//TODO принимать в user.New по interface` — значит, проблема осознаётся.

**Рекомендация:** вернуть интерфейсы, но объявлять их **на стороне потребителя** (в пакете `service`), а не в `repository` — это идиоматический Go-подход.

---

### Непоследовательное именование пакетов

| Текущее имя | Проблема | Идиоматично |
|---|---|---|
| `serviceT` | PascalCase-хвост, нечитаемо | `testsvc` / `testservice` |
| `loginservice` | OK | `loginsvc` / `loginservice` |
| `setupT.go`, `setupL.go` | Суффиксы-одиночки | `setup.go` в каждом пакете |

В пределах одного модуля — два разных стиля. Go-конвенция: короткие lowercase-имена одним словом.

---

## 🟡 Мелкие проблемы и технический долг

| Место | Проблема |
|---|---|
| `internal/repository/storage/test/methods.go:38` | `fmt.Println(testfull)` — дебажный вывод, оставленный в проде |
| `internal/repository/storage/test/methods.go:46` | `s.db.Query` вместо `QueryContext` — `ctx` передан, но игнорируется |
| `internal/repository/storage/login/methods.go:20` | То же: `ctx` передан, но используется `db.Query` |
| `cmd/main.go:44` | `http_transport.New(cfg.Server, *handler1, *handler2)` — разыменование указателя ради копии структуры с указателем внутри; передавать надо по указателю |
| `internal/transport/http_transport/handler/loginHandler.go:11-14` | Дубль DTO: `requestBody` ≡ `model.LoginRequest` |
| `internal/transport/http_transport/handler/loginHandler.go:47` | Опечатка в JSON-ключе: `"responce"` → `"response"` |
| `internal/transport/http_transport/handler/loginHandler.go:40` | 400 Bad Request на «user not found» — должно быть 401; на DB-ошибке — 500 |
| `internal/transport/http_transport/handler/loginHandler.go:48` | Голый `return` в конце функции — лишний |
| `internal/transport/http_transport/responses.go` | Весь файл — мёртвый код (`errInvalidUserId`, `errResp`, `newErrResp` нигде не используются) |
| `internal/model/errors.go` | `ErrUserNotFound` объявлен и не используется |
| `internal/service/serviceT/service.go:14,22,23` | Опечатка `avaiilableTests` |
| `internal/config/config.go:45` | Текст ошибки «using system environment variables» не соответствует поведению (нет fallback) |
| `.gitignore` | `go.sum` в игноре — **антипаттерн**, lock-файл должен коммититься |
| `cmd/main.go:32` | Устаревший TODO: `user.New` больше не существует, теперь `test.New` |

---

## Сводная таблица по степени риска

| Уровень | Проблема | Эффект |
|---|---|---|
| 🔴 | Утечка `rows.Close()` | Сервер встанет после ~25 запросов |
| 🔴 | `Password int` | Невозможно логиниться с нормальным паролем, нет хеша |
| 🔴 | Windows-путь в `config.Load` | Приложение не запускается на macOS/Linux |
| 🔴 | Проглоченная ошибка в `loginService.SignIn` | DB-ошибка маскируется под «user not found» |
| 🔴 | Игнор ошибки + N+1 в `GetAvailableTests` | Скрытые баги + деградация производительности |
| 🟠 | Удалены интерфейсы | Невозможно писать unit-тесты, жёсткая связность слоёв |
| 🟠 | Непоследовательные имена пакетов (`serviceT`, `setupT`, `setupL`) | Нарушены Go-конвенции, затрудняет чтение |
| 🟡 | `fmt.Println` в репозитории | Мусор в логах |
| 🟡 | Игнор `ctx` (`Query` вместо `QueryContext`) | Нет отмены по таймауту |
| 🟡 | Мёртвый код (`responses.go`, `errors.go`) | Технический долг |
| 🟡 | Опечатки (`responce`, `avaiilableTests`, `loginZ`, `passwordZ`) | API-контракт и читаемость |
| 🟡 | `go.sum` в `.gitignore` | Сборка невоспроизводима между машинами |

---

## Приоритетный план исправления

1. **Починить `rows.Close()`** во всех репозиториях — иначе прод упадёт.
2. **Сменить `Password` на `string` + bcrypt** — сейчас авторизация по сути сломана.
3. **Убрать хардкод пути к `.env`** — без этого проект не стартует на не-Windows.
4. **Вернуть интерфейсы** для `TestRepository` и `LoginRepository` (объявить в `service/`), чтобы открыть дорогу юнит-тестам.
5. **Исправить обработку ошибок** в `loginService.SignIn` и `serviceT.GetAvailableTests`.
6. **Переписать `GetAvailableTests`** одним JOIN-запросом на уровне репозитория (как было раньше).
7. **Использовать `QueryContext`** везде, где уже передаётся `ctx`.
8. **Почистить мёртвый код** (`responses.go`) и удалить дебажные `Println`.
9. **Привести имена пакетов к Go-стилю** (`serviceT` → `testservice`, `setupT.go`/`setupL.go` → `setup.go`).
10. **Убрать `go.sum` из `.gitignore`** и закоммитить lock-файл.
