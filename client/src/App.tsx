import React, {useEffect, useState} from "react";
import {CodeResponse, useGoogleLogin} from "@react-oauth/google";
import "./App.css";
import axios from "axios";

interface User {
    name: string;
    email: string;
    picture: string;
}

const App: React.FC = () => {
    const [user, setUser] = useState<User | null>(null);

    const googleLogin = useGoogleLogin({
        flow: "auth-code",
        onSuccess: async (codeResponse: CodeResponse) => {
            const resp = await axios.post("http://localhost:8000/api/login/google",
                {code: codeResponse.code},
                {withCredentials: true}
            );
            setUser({
                name: resp.data.name,
                email: resp.data.email,
                picture: resp.data.picture,
            })
        },
        onError: (errorResponse: any) => console.log(errorResponse),
    });

    const logOut = async () => {
        try {
            await axios.get("http://localhost:8000/api/logout",{withCredentials: true});
            setUser(null);
            console.log("User logged out successfully");
        } catch (error) {
            console.error("Error logging out:", error);
        }
    };

    useEffect(() => {
        const checkCookie = async () => {
            try {
                const resp = await axios.get(
                    "http://localhost:8000/api/user",
                    {withCredentials: true}
                );
                setUser({
                    name: resp.data.name,
                    email: resp.data.email,
                    picture: resp.data.picture,
                })
            } catch (error) {
                console.log("Cookie is not valid");
            }
        };
        checkCookie();
    }, []);

    return (
        <div className="container">
            <h2>Oauth2 example (Golang & Reactjs) </h2>
            <br/>
            <br/>
            {user ? (
                <div className="user-info">
                    <img src={user.picture} alt="profile" id="user-image"/>
                    <h3>User Logged in</h3>
                    <p>Name: {user.name}</p>
                    <p>Email Address: {user.email}</p>
                    <br/>
                    <br/>
                    <button onClick={logOut} className="button" id="logout-button">Log out</button>
                </div>
            ) : (
                <button onClick={googleLogin} className="button" id="login-button">Sign in with Google 🚀</button>
            )}
        </div>
    );
}

export default App;