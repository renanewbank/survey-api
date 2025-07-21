
# Survey API

API REST desenvolvida em Go para gerenciamento de perguntas de pesquisa (survey questions).

---

##  Funcionalidades Implementadas

* CRUD completo da entidade `Question`:

  * Criar, listar, buscar por ID, atualizar e deletar perguntas.
* Validação contra schema JSON travado (`api/question_schema.json`).
* Versionamento automático de perguntas.
* Armazenamento em memória (thread-safe).
* Estrutura modular preparada para extensão (ex: auditoria, publicação).

---

##  Estrutura do Projeto

```
survey-api/
├── cmd/server/               # Entrada principal da aplicação
├── internal/
│   ├── models/               # Definições de structs (ex: Question)
│   ├── handlers/             # Handlers HTTP
│   ├── repository/           # Persistência em memória
├── pkg/jsonschema/           # Validação de schema JSON
├── api/question_schema.json  # Schema JSON congelado
├── go.mod / go.sum           # Gerenciamento de dependências Go
```

---

##  Como executar localmente

1. Clone o repositório:

```bash
git clone https://github.com/seuusuario/survey-api.git
cd survey-api
```

2. Execute o servidor:

```bash
go run cmd/server/main.go
```

A API ficará disponível em:
**[http://localhost:8080](http://localhost:8080)**

---

##  Endpoints disponíveis

### POST `/questions`

Cria uma nova pergunta.

### GET `/questions`

Lista todas as perguntas.

### GET `/questions/{id}`

Busca uma pergunta específica pelo ID.

### PUT `/questions/{id}`

Atualiza uma pergunta existente.

### DELETE `/questions/{id}`

Remove uma pergunta do sistema.

---

##  Validação de schema

* O JSON enviado para criação/edição é validado contra o schema fixo em `api/question_schema.json`.
* Campos como `id`, `version`, `created_at`, `updated_at` são controlados pelo backend.

---

##  Observações

* A persistência atual é em memória (não persistente entre execuções).
* A funcionalidade de publicação/despublicação e logs de auditoria será implementada posteriormente.
* O projeto segue boas práticas de separação de responsabilidades e modularização para facilitar manutenção e expansão.

