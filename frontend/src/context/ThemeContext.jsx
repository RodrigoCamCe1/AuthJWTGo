import React, { createContext, useState, useContext } from 'react';

// 1. Crear el contexto y EXPORTARLO (para que useTheme lo encuentre)
export const ThemeContext = createContext();

// 2. Crear el Provider y EXPORTARLO (para que App.jsx lo encuentre)
export const ThemeProvider = ({ children }) => {
    const [theme, setTheme] = useState('dark'); // Por defecto 'dark'

    const toggleTheme = () => {
        setTheme((prevTheme) => (prevTheme === 'light' ? 'dark' : 'light'));
    };

    return (
        <ThemeContext.Provider value={{ theme, toggleTheme }}>
            {children}
        </ThemeContext.Provider>
    );
};

// 3. Hook personalizado (opcional, por si lo usas en el mismo archivo)
export const useTheme = () => useContext(ThemeContext);