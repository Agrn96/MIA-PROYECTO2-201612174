import './App.css';
import {useState} from 'react';
import { Button } from './components/Button';
import { saveAs } from 'file-saver';
import { Input } from './components/Input';
import { userApi } from './Api/userApi';

function App() {
  const [myValue, setMyValue] = useState('');
  const [myValue2, setMyValue2] = useState('');
  

    const handleClick = () => {
   
    console.log(myValue);
    userApi.post('http://ec2-3-137-107-30.us-east-2.compute.amazonaws.com:8080/interpretar', { entrada: myValue }
    //userApi.post('http://localhost:8080/interpretar', { entrada: myValue }
    )
    .then(({data}) => {
      const {entrada} = data
      console.log(entrada)
      setMyValue2(entrada)
    });
     
      

  }
  const createFile = () => {
    const blob = new Blob([myValue], { type: 'text/plain;charset=utf-8' })
    saveAs(blob, 'myFile.txt')
  }

  const readFile = (e) => {

    const file = e.target.files[0];
    if (!file) return;

    const fileReader = new FileReader();

    fileReader.readAsText(file);

    fileReader.onload = () => {
      console.log(fileReader.result);
      setMyValue(fileReader.result);
    }

    fileReader.onerror = () => {
      console.log(fileReader.error);
    }

  }
  return (
    <div className="App">

      <header className="App-header">

        <h1>ADMINISTRADOR DE ARCHIVOS</h1>
        <form>
          <textarea cols="75"
            rows="35"
            placeholder="Ingreser codigo aqui"
            value={myValue}
            onChange={(e) => setMyValue(e.target.value)}
          >
          </textarea>
          <Button onClick={handleClick}
            type="button"
            buttonStyle="btn--primary--solid"
            buttonSize="btn--medium"
          >Interpretar
          </Button>
        </form>
        <Button onClick={createFile}
          type="submit"
          buttonStyle="btn--primary--solid"
          buttonSize="btn--medium"
        >Nuevo archivo
        </Button>
        <Input
          type="file"
          onChange={readFile}
          inputStyle={'imp--primary--solid'}
          inputSize={'imp--small'}
        >
          Cargar archivo
        </Input>
        <textarea cols="60"
          rows="10"
          placeholder="Consola"
          value={myValue2}
          onChange={(e) => setMyValue2(e.target.value)}
        >

        </textarea>
    

     </header>
     
     

    
    </div>
  );
}

export default App;
