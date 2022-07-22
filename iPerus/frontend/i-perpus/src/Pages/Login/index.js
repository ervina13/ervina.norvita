// import React from "react";
// import { Link } from "react-router-dom";

// const Login = () => {
//   return (
//     <>
//       <Link to="/home">Login</Link>
//     </>
//   );
// };

// export default Login;

import React, { useState } from "react";

function LoginForm({ Login, error }) {
  const [details, setDetails] = useState({ name: "", password: "" });

  const submit = async (e) => {
    e.preventDefault();
    Login(details);
  };

  // const resp = await fetch('http://localhost:2213/api/login', {
  //     method: 'POST',
  //     headers: { 'Content-Type': 'application/json' },
  //     credentials: 'include',
  //     body: JSON.stringify({
  //         name,
  //         password,
  //     })
  // })

  // let code = resp.status
  // let content = await resp.json()

  // const user = await fetch('http://localhost:2213/', {
  //     method: 'GET',
  //     headers: { 'Content-Type': 'application/json' },
  //     credentials: 'include',
  // })

  // const totalUser = await user.json()
  // const allUser = totalUser.data

  return (
    <div className="App">
      <form onSubmit={submit}>
        <div className="form-inner">
          <h2>Login</h2>
          {error != "" ? <div className="error">{error}</div> : <div></div>}
          <div className="form-group">
            <label htmlFor="name">Name:</label>
            <input
              type="text"
              name="name"
              id="name"
              onChange={(e) => setDetails({ ...details, name: e.target.value })}
              value={details.name}
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              name="password"
              id="password"
              onChange={(e) =>
                setDetails({ ...details, password: e.target.value })
              }
              value={details.password}
            />
          </div>
          <input className="login-btn" type="submit" value="LOGIN" />
          <div>
            <br />

            <input className="regis-btn" type="submit" value="REGISTRASI" />
          </div>
        </div>
      </form>
    </div>
  );
}

export default LoginForm;
