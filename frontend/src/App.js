import './App.css';
import AddressList from "./components/AddressList";
import AddAddress from "./components/AddAddress";

function App() {
  return (
    <div className='App-header'>
      <h1>Bitcoin Wallet</h1>
      <AddAddress />
      <AddressList />
    </div>
  );
}

export default App;
