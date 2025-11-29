import React from 'react';
import Header from '../components/Header';
import Sidebar from '../components/Sidebar';
import { Outlet } from 'react-router-dom';

const MainLayout = () => {
    return (
        <div className="main-layout" style={{ display: 'flex', minHeight: '100vh' }}>
            {/* Barra Lateral */}
            <Sidebar />

            <div style={{ flex: 1, display: 'flex', flexDirection: 'column' }}>
                {/* Encabezado Superior */}
                <Header />

                {/* Contenido de la Página (Aquí se cargan Home, Employees, etc.) */}
                <main style={{ padding: '20px', backgroundColor: '#f4f6f8', flex: 1 }}>
                    <Outlet />
                </main>
            </div>
        </div>
    );
};

export default MainLayout;