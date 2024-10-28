import React, { useEffect, useState } from 'react';
import { Input, Stack } from '@chakra-ui/react';

const CategoriesPage: React.FC = () => {

    useEffect(() => {
        
    }, []);

    return (
        <Stack gap="4">
            <Input placeholder="size (xs)" size="xs" />
            <Input placeholder="size (sm)" size="sm" />
            <Input placeholder="size (md)" size="md" />
            <Input placeholder="size (lg)" size="lg" />
        </Stack>
    );
};

export default CategoriesPage;
