# OAuth2 Authentication with Golang and React

This project demonstrates OAuth2 authentication using Golang for the backend API and React for the frontend application.

## Getting Started

### Prerequisites:
- Golang (>= 1.22.5) installed: [Download Golang](https://golang.org/dl/)
- Node.js and npm (or yarn) installed: [Download Node.js](https://nodejs.org/en/)
- Provider account (Google) to register your application

### Installation:
1. Clone this repository:
   ```bash
   git clone https://github.com/iarsham/oauth2-example.git
   ```


2. Navigate to the project directory:
   ```bash
   cd oauth2-example
   ```

3. Install Go dependencies:
   ```bash
   go mod download
   ```

4. Install React dependencies:
   ```bash
   cd client
   npm install
   ```

### Configuration

1. **Provider Registration:**
    - Visit your chosen provider's developer console (e.g., Google Cloud Platform for Google OAuth2).
    - Create a new project or select an existing one.
    - Enable the OAuth2 API for your project.
    - Create OAuth credentials, specifying:
        - Authorized JavaScript origins: The URL of your React application (e.g., http://localhost:3000).
        - Authorized redirect URIs: A redirect URI where the provider will send the authorization code after successful login (e.g., http://localhost:3000).
    - Copy the Client ID and Client Secret provided by the provider.


2. **Add yaml properties in configs folder (Golang):**
    - Fill the variables with your specific value(Postgres user, password, ...):
    - Also fill Client-ID in client/src/index.tsx for integrate with google and backend



### Running the Application

1. Start the Golang backend server:
   ```bash
   go run ./cmd/web
   ```

2. Start the React development server:
   ```bash
   cd client
   npm start
   ```

3. Access your React application in a web browser (usually http://localhost:3000).
