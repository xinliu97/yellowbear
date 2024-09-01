// GeneralLayout.js
import React from 'react';
import {Divider, Layout, Menu, Space} from 'antd';
import {NavLink, useLocation, useNavigate} from "react-router-dom";

const { Header, Content, Footer } = Layout;

const GeneralLayout = ({ children }) => {
    const navigate = useNavigate(); // 获取导航函数
    const location = useLocation(); // 获取当前路径


    const menuItems = [
        {
            key: '/',
            label: '主页',
            onClick: () => navigate('/'),
        },
        {
            key: '/quiz',
            label: '测试',
            onClick: () => navigate('/quiz'),
        },
        {
            key: '/settings',
            label: '设置',
            onClick: () => navigate('/settings'),
        },
    ];
    return (
        <Layout style={{ minHeight: '100vh' }}> {/* 设置最小高度为100%视口高度 */}
            <Header style={{ backgroundColor: 'white', display: 'flex', alignItems: 'left', fontFamily: 'Arial' }}>
                {/* Logo 部分 */}
                <div style={{ display: 'flex', alignItems: 'center', marginRight: '10px' }}>
                    <img
                        src="/bearlogo.png"
                        alt="Logo"
                        style={{ height: '40px' }}
                    />
                </div>
                <Menu
                    style={{ backgroundColor: 'white' }}
                    theme='light'
                    mode="horizontal"
                    items={menuItems}
                    selectedKeys={[location.pathname]}
                />
            </Header>

            {/* 分割线 */}
            <Divider style={{ margin: 0 }} />

            <Content style={{ padding: '20px', paddingBottom: '60px' }}> {/* 给内容区域添加底部填充 */}
                {children}
            </Content>

            {/* 分割线 */}
            <Divider style={{ margin: 0 }} />

            <Footer style={{ textAlign: 'center', position: 'fixed', left: 0, bottom: 0, width: '100%', backgroundColor: '#f0f2f5' }}>
                My Application ©2024 Created by YellowBear
            </Footer>
        </Layout>
    );
}

export default GeneralLayout;
