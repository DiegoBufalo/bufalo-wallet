import { RouterProvider } from 'react-router-dom';
import { Provider } from './components/ui/provider';
import { router } from './routes';

export default function App() {
    return (
        <Provider>
            <RouterProvider router={router} />
        </Provider>
    );
}
