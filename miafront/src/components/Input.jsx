import React from "react";
import "./input.css";

    const STYLE = [
        "imp--primary--solid",
    ]

    const SIZES = ["imp--medium", "imp--large", "imp--small"]

export const Input = ({ 
type,
onChange, 
inputStyle,
inputSize,
}) => {

    const checkInputStyle = STYLE.includes(inputStyle)
    ? inputStyle:
    STYLE[0]

    const checkInputSize = SIZES.includes(inputSize)
    ? inputSize
    : SIZES[0]

    return(
        <input className={`imp ${checkInputStyle} ${checkInputSize}`}
        type={type}
        onChange={onChange}
        
        />
           
       
    )
}