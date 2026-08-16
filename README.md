# CRUD-Go
Meu primeiro CRUD em GOLang

## Requisitos
- Docker + Docker Compose

## Como rodar

1. Copie o arquivo de ambiente (opcional, valores padrão já funcionam):
   ```bash
   cp .env.example .env
   ```

2. Suba a aplicação e o banco de dados PostgreSQL:
   ```bash
   docker compose up -d --build
   ```

3. A API estará disponível em `http://localhost:2611` e o PostgreSQL em `localhost:5432`.

A tabela `users` é criada automaticamente pelo `db/init.sql` na primeira subida.

## Endpoints

| Método | Rota                     | Descrição                   |
|--------|--------------------------|-----------------------------|
| POST   | `/createUser`            | Cria um usuário             |
| GET    | `/listAllUsers`          | Busca por todos os usuarios |
| GET    | `/getUserByID/:userId`   | Busca por ID (UUID)         |
| GET    | `/getUserByEmail/:email` | Busca por email             |
| PUT    | `/updateUser?userId=:id` | Atualiza usuário            |
| DELETE | `/deleteUser/:UserId`    | Deleta usuário              |

Exemplo de corpo para POST/PUT:
```json
{
  "name": "Pedro Bett",
  "age": 25,
  "email": "pedro@example.com",
  "password": "secret!123"
}
```

## Variáveis de ambiente

| Variável     | Padrão   | Descrição              |
|--------------|----------|------------------------|
| `DB_USER`    | `crud`   | Usuário do PostgreSQL  |
| `DB_PASSWORD`| `crud`   | Senha do PostgreSQL    |
| `DB_HOST`    | `localhost` | Host do banco        |
| `DB_PORT`    | `5432`   | Porta do banco         |
| `DB_NAME`    | `crud`   | Nome do banco          |

Para rodar localmente (sem Docker), use `DB_HOST=localhost` e um PostgreSQL de sua escolha.
