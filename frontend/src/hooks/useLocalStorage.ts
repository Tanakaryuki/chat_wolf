import { useEffect, useState } from "react";

export const useLocalStorage = (key: string, initialValue: string) => {
	const [storedValue, setStoredValue] = useState(() => {
		try {
			const item = localStorage.getItem(key);
			return item ? JSON.parse(item) : initialValue;
		} catch (error) {
			console.log(error);
			return initialValue;
		}
	});

	useEffect(() => {
		try {
			const serializedValue = JSON.stringify(storedValue);
			localStorage.setItem(key, serializedValue);
		} catch (error) {
			console.log(error);
		}
	}, [key, storedValue]);

	return [storedValue, setStoredValue];
};
