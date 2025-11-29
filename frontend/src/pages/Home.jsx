import React from 'react';
import './Home.css'; // Esto usará el archivo CSS que ya tienes ahí

const Home = () => {
    return (
        <div className="home-container" style={{ padding: '40px', textAlign: 'center' }}>
            <h1>🏠 Bienvenido al Sistema EMS</h1>
            <p>¡Has iniciado sesión correctamente!</p>
            <div style={{ marginTop: '20px', padding: '20px', border: '1px solid #ccc', borderRadius: '8px' }}>
                <h3>Panel de Control</h3>
                <p>Selecciona una opción del menú para comenzar.</p>
            </div>
        </div>
    );
};

export default Home;