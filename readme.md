# QR Code Scanner for Stadium
## Project Structure: 

```shell script
stadium/
|-- go-backend/
|   |-- main.go
|   |-- api/
|   |   |-- handlers/
|   |   |   |-- user_handler.go
|   |   |   |-- qr_code_handler.go
|   |   |-- routes.go
|   |   |-- middleware.go
|   |-- database/
|   |   |-- postgres/
|   |   |   |-- postgres.go
|   |   |-- models/
|   |   |   |-- user.go
|   |   |   |-- qr_code.go
|   |   |-- sql/
|   |   |   |-- schema/
|   |   |   |   |-- 20231125000000_CreateUsersTable.sql
|   |   |   |   |-- 20231125000001_CreateQRCodesTable.sql
|   |   |   |-- queries/
|   |   |   |   |-- user.sql
|   |   |   |   |-- qr_code.sql
|   |   |-- sqlc.yaml
|   |-- utils/
|   |   |-- utils.go
|-- python-hardware/
|   |-- main.py
|   |-- hardware/
|   |   |-- qr_code_scanner.py
|   |-- api_client/
|   |   |-- api_client.py
|-- postgres/
|   |-- init.sql
|-- web-frontend/
|   |-- public/
|   |-- src/
|   |   |-- components/
|   |   |   |-- UserRegistration.js
|   |   |   |-- QRCodeScanner.js
|   |   |-- services/
|   |   |   |-- apiService.js
|   |   |-- App.js
|   |   |-- index.js
|-- .gitignore
|-- README.md
|-- go.mod
|-- go.sum
|-- requirements.txt
|-- package.json
|-- yarn.lock
|-- .env
|-- docker-compose.yml
```

## How To Make Changes:
1. Ensure you're on the main branch:
```shell script
git checkout main
```
2. Pull the latest changes from the remote main branch:
```shell script
git pull origin main
```
3. Create a new branch for the new feature/change:
```shell script
git checkout -b feature/new-feature
```
  This command creates and checks out a new branch named feature/new-feature. Replace new-feature with a descriptive name related to the feature or task you're working on.

4. Stage and commit your changes:
```shell script
git add .
git commit -m "Implementing new feature"
```
5. Push your feature branch to the remote repository:
```shell script
git push origin feature/new-feature
```
6. Create a Pull Request (PR) on GitHub or your Git hosting platform:

      Navigate to your repository on GitHub or your Git hosting platform.
      Find the option to create a new pull request.
      Select your feature branch (feature/new-feature) as the source branch and master as the target branch.
      Provide a title and description for your pull request, summarizing the changes made.
      Submit the pull request.

8. Code Review:

     Review the code changes.

9. Merge the Pull Request:

      After the code review process is successful, you can merge the pull request.
      Click the "Merge" button on your pull request on GitHub or your Git hosting platform.

10. Update the main branch:
```shell script
git checkout main
git pull origin main
```
