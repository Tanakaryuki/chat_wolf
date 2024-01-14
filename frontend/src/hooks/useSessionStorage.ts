import { useEffect, useState } from "react";

export const useSessionStorage = (key: string, initialValue: string) => {
	const [storedValue, setStoredValue] = useState(() => {
		try {
			const item = sessionStorage.getItem(key);
			return item ? JSON.parse(item) : initialValue;
		} catch (error) {
			console.log(error);
			return initialValue;
		}
	});

	useEffect(() => {
		try {
			const serializedValue = JSON.stringify(storedValue);
			sessionStorage.setItem(key, serializedValue);
		} catch (error) {
			console.log(error);
		}
	}, [key, storedValue]);

	return [storedValue, setStoredValue];
};
