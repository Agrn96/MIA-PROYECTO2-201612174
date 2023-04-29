import React from "react";
import "./button.css";

const STYLE = [
    "btn--primary--solid", 
    "btn--danger--solid", 
    "btn--warning--solid", 
    "btn--success--solid",
    "btn--primary--outline",
    "btn--danger--outline",
    "btn--warning--outline",
    "btn--success--outline", 
]

const SIZES = ["btn--medium", "btn--large", "btn--small"]

export const Button = ({ children, 
    type, 
    onClick, 
    buttonStyle, 
    buttonSize, 
    }) => {

        const checkButtonStyle = STYLE.includes(buttonStyle) 
        ? buttonStyle : 
        STYLE[0] 

        const checkButtonSize = SIZES.includes(buttonSize) 
        ? buttonSize 
        : SIZES[0]
        
    
        return (
            <button className={`btn ${checkButtonStyle} ${checkButtonSize}` } 
            onClick={onClick} 
            type={type}
            >
                {children}
            </button>
        )
        }