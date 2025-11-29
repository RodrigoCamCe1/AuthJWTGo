import React from 'react';
import { Link } from 'react-router-dom';

const Sidebar = () => {
    return (
        <div className="w-64 bg-gray-900 text-white min-h-screen p-4">
            <h2 className="text-2xl font-bold mb-6">EMS System</h2>
            <ul>
                <li className="mb-2"><Link to="/" className="hover:text-gray-300">Inicio</Link></li>
                <li className="mb-2"><Link to="/employees" className="hover:text-gray-300">Empleados</Link></li>
            </ul>
        </div>
    );
};
export default Sidebar;