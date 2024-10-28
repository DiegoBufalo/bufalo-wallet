import CategoriesPage from "../pages/categories";
import { createBrowserRouter } from "react-router-dom";

export const router = createBrowserRouter([
    {
        path: "/categories",
        element: (
            <CategoriesPage />
        ),
    },
]);