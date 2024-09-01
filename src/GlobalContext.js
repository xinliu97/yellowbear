import React, { createContext } from 'react';
import {ConfigProvider} from "antd";

const GlobalContext = createContext();

export const GlobalProvider = ({ children }) => {
    // const [globalVariable, setGlobalVariable] = useState('This is a global variable');
    //set global color

    return (
        <ConfigProvider
            theme={{
                token: {
                    colorPrimary: '#00b96b', // 主色，所有按钮和组件的主要颜色
                    borderRadius: 2,          // 全局的圆角半径
                    colorBgContainer: '#fadb14', // 容器背景色
                },
            }}
        >
            {children}
        </ConfigProvider>
    );
};

export default GlobalContext;
