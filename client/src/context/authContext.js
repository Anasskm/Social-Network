import { createContext, useEffect, useState } from "react";
import Profile from "../pages/profile/Profile";
import axios from "axios";












export const AuthContext = createContext();



export const AuthContextProvider = ({ children }) => {
    const [currentUser, setCurrentUser] = useState(
        JSON.parse(localStorage.getItem("user")) || null
    );


    const login = async (inputs) => {
        const res = await axios.post("http://localhost:8080/login", inputs, {
            withCredentials: true,
        });
        console.log(res)
        setCurrentUser(res.data)
    };







    useEffect(() => {
        localStorage.setItem("user", JSON.stringify(currentUser));
    }, [currentUser])
    return (
        <AuthContext.Provider value={{ currentUser, login }} >
            {children}
        </AuthContext.Provider>
    )
}






