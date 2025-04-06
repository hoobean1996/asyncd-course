import { RelayEnvironmentProvider } from "react-relay";
import { RelayEnvironment } from "./RelayEnvironment";
import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Outlet, Route, Routes } from "react-router-dom";
import App from "./App";
import AdminPage from "./pages/AdminPage";
import TasksPage from "./pages/TasksPage";
import HandlersPage from "./pages/HandlersPage";


const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

// /admin ()
//  /admin/tasks -> Outlet 就会被替换为 tasks's element
//  /admin/handlers -> Outlet 就被会替换为 handers's elemnet


// Tailwind --> CSS 简化了CSS库
// <div styles={{width: 100, height: 100px}}> </div>
// <div className="w1 w2..."> 

// 使用组件库好处： 
//  1. 容易上手，不容易去定制化 AntD 
// 使用CSS库好处：
//  2. 在Gen AI时代，使用tailwind 实现页面 ->  抽象成自己的组件 自己的风格
root.render(
  <RelayEnvironmentProvider environment={RelayEnvironment}>
    <React.StrictMode>
      <BrowserRouter>
        <Routes>
          <Route path="/admin" element={<AdminPage />}>
            <Route path="tasks" element={<TasksPage />} />
            <Route path="handlers" element={<HandlersPage />} />
            <Route path="analysts" element={<div>analysts</div>} />
            <Route path="settings" element={<div>settings</div>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </React.StrictMode>
  </RelayEnvironmentProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
