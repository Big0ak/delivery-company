import { createContext } from "react";

interface AppContextProps {
  isLoggedIn: boolean;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>> | (() => void);
}

const AppContext = createContext<AppContextProps>({
  isLoggedIn: false,
  setIsLoggedIn: () => {},
});

export { AppContext };