import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import {GoogleOAuthProvider} from "@react-oauth/google"

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);

root.render(
    <GoogleOAuthProvider clientId='167649986158-jeq40b37r3tpc7i1m084befpktf7gra9.apps.googleusercontent.com'>
        <React.StrictMode>
            <App/>
        </React.StrictMode>
    </GoogleOAuthProvider>
);

