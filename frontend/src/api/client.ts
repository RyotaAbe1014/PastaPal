export const ApiClient = () => {
	const settings = {
		BaseUrl: import.meta.env.VITE_ENV_API_URL,
	};
	const request = async <RequestType = undefined, ResponseType = unknown>(
		path: string,
		method: string,
		params?: RequestType,
	): Promise<ResponseType> => {
		const url = `${settings.BaseUrl}${path}`;
		const options: RequestInit = {
			method,
			credentials: "include",
			mode: "cors",
			headers: {
				"Content-Type": "application/json",
			},
			body: params ? JSON.stringify(params) : undefined,
		};
		const response = await fetch(url, options);
		return handleResponse<ResponseType>(response);
	};

	const handleResponse = async <ResponseType>(
		response: Response,
	): Promise<ResponseType> => {
		if (!response.ok) {
			// TODO: ここのエラーハンドリングは適切なものに変更する
			throw new Error("fetch failed");
		}
		return response.json() as Promise<ResponseType>;
	};

	const Get = async <RequestType = undefined, ResponseType = unknown>(
		path: string,
		params?: RequestType,
	): Promise<ResponseType> => {
		// paramsをクエリパラメータに変換
		const queryString = params ? requestToUrlSearch(params).toString() : "";
		// クエリパラメータがあればURLに追加
		const requestUrl = queryString ? `${path}?${queryString}` : path;

		return request(requestUrl, "GET", params);
	};
	const Post = <RequestType = undefined, ResponseType = unknown>(
		path: string,
		params?: RequestType,
	): Promise<ResponseType> => request(path, "POST", params);
	const Put = <RequestType = undefined, ResponseType = unknown>(
		path: string,
		params?: RequestType,
	): Promise<ResponseType> => request(path, "PUT", params);
	const Delete = <RequestType = undefined, ResponseType = unknown>(
		path: string,
		params?: RequestType,
	): Promise<ResponseType> => request(path, "DELETE", params);

	return {
		Get,
		Post,
		Put,
		Delete,
	};
};

const requestToUrlSearch = <
	RequestType extends Record<string, unknown> | Record<string, unknown>[],
>(
	request: RequestType,
) => {
	const searchParams = new URLSearchParams();
	// biome-ignore lint: lint/complexity/noForEach
	Object.entries(request).forEach(([key, value]) => {
		if (Array.isArray(value)) {
			// biome-ignore lint: lint/complexity/noForEach
			value.forEach((v) => searchParams.append(`${key}[]`, String(v)));
		} else {
			searchParams.append(key, String(value));
		}
	});
	return searchParams;
};
