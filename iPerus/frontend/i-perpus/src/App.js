import "./index.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./Pages/Users/Content/Homepage";
import SignInPage from "./Pages/signin";
import LoginForm from "./Pages/Login";
import KaryaTulis from "./Pages/Users/Content/KaryaTulis";
import KaryaTulisDetail from "./Pages/Users/Content/KaryaTulis/KaryaTulisDetail";
import Contact from "./Pages/Users/Content/Contact";
import Profile from "./Pages/Users/Content/Profile";
import React, { useState } from "react";

function App() {
  const user = {
    name: "Abdi",
    password: "1010",
  };

  const [login, setLogin] = useState({ name: "", password: "" });
  const [error, setError] = useState("");

  const Login = (details) => {
    console.log(details);

    if (details.name === user.name && details.password === user.password) {
      console.log("Login success");
      setLogin({
        name: details.name,
        password: details.password,
      });
    } else {
      console.log("Login failed");
      setError("Login failed!");
    }
  }; // end of Login

  const Logout = () => {
    setLogin({
      name: "",
      password: "",
    });
  }; // end of Logout

  return (
    <>
      {login.name != "" ? (
        <Router>
          <Routes>
            <Route path="/signin" element={<SignInPage />} exact />
            <Route path="/" element={<Home />} exact />
            <Route path="/home" element={<Home />} exact />
            <Route path="/karya-tulis" element={<KaryaTulis />} exact />
            <Route
              path="/karya-tulis/:id"
              element={<KaryaTulisDetail />}
              exact
            />
            <Route path="/contact" element={<Contact />} exact />
            <Route path="/profile" element={<Profile />} exact />
          </Routes>
        </Router>
      ) : (
        <LoginForm Login={Login} error={error} />
      )}
    </>
  );
}

export default App;
