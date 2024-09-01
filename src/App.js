import './App.css';
import GeneralLayout  from "./GeneralLayout";
import RoutesConfig from "./RoutesConfig";
// import {GlobalProvider} from "./GlobalContext";
import {Button, ConfigProvider} from "antd";
import {Header} from "antd/es/layout/layout";
import {BrowserRouter} from "react-router-dom";

function App() {

    return (
        <ConfigProvider
            theme={{
                token: {
                    // Seed Token，影响范围大
                    colorPrimary: '#d4b106',
                    borderRadius: 2,

                    // 派生变量，影响范围小
                    colorBgContainer: '#f6ffed',
                },
            }}
        >
            <BrowserRouter>
                <GeneralLayout>
                    <RoutesConfig />
                </GeneralLayout>
            </BrowserRouter>
        </ConfigProvider>
  );
}

export default App;
