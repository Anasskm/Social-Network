import LeftBar from "./components/leftBar/LeftBar";
import Navbar from "./components/navbar/Navbar";
import RightBar from "./components/rightBar/RightBar";
import Login from "./pages/login/Login";
import Register from "./pages/register/Register";
import Profile from "./pages/profile/Profile"
import Home from "./pages/home/Home";
import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Outlet,
  Navigate,
} from "react-router-dom";
import { useContext, useState, useEffect } from "react";
import { AuthContext } from "./context/authContext";
import { QueryClient, QueryClientProvider } from 'react-query';
import Loading from "./pages/loading/Loading";








function App() {
  const { currentUser } = useContext(AuthContext);
  const queryClient = new QueryClient();








  const Layout = () => {
    return (


      <QueryClientProvider client={queryClient}>
        <div style={{ backgroundColor: "#fafafa" }}>
          <Navbar />
          <div style={{ display: "flex" }}>
            <LeftBar />
            <div style={{ flex: 6 }}>
              <Outlet />
            </div>

            <RightBar />
          </div>
        </div>

      </QueryClientProvider>

    )

  };
  const ProtectedRoute = ({ children }) => {
    if (!currentUser) {
      return <Navigate to="/login" />
    }
    return children
  };


  const router = createBrowserRouter([
    {
      path: "/",
      element: (<ProtectedRoute>
        <Layout />
      </ProtectedRoute>
      ),
      children: [{
        path: "/profile/:id",
        element: <Profile />,
      },
      {
        path: "/",
        element: <Home />,
      }
      ]
    },
    {
      path: "/login",
      element: <Login />,
    },
    {
      path: "/register",
      element: <Register />,
    }, {
      path: "/l",
      element: (<ProtectedRoute>
        <Loading />
      </ProtectedRoute>)

    }
  ]);
  return <div>
    <RouterProvider router={router} />
  </div>;
}

export default App;
