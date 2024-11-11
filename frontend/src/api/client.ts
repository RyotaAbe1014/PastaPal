export const ApiClient = () => {
	const settings = {
		withCredentials: true,
		headers: {
			"Content-Type": "application/json",
		},
		BaseUrl: process.env.REACT_APP_API_URL,
	};
	const request = async <RequestType, ResponseType>(
		path: string,
		method: string,
		params?: RequestType,
	): Promise<ResponseType> => {
		const url = new URL(path, settings.BaseUrl);
		const options: RequestInit = {
			...settings,
			method,
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

	const Get = async <RequestType, ResponseType>(
		path: string,
		params?: RequestType,
	): Promise<ResponseType> => {
		// paramsをクエリパラメータに変換
		const queryString = params ? requestToUrlSearch(params).toString() : "";
		// クエリパラメータがあればURLに追加
		const requestUrl = queryString ? `${path}?${queryString}` : path;
		const url = new URL(requestUrl, settings.BaseUrl).toString();

		return request(url, "GET", params);
	};
	const Post = <RequestType, ResponseType>(
		path: string,
		params: RequestType,
	): Promise<ResponseType> => request(path, "POST", params);
	const Put = <RequestType, ResponseType>(
		path: string,
		params: RequestType,
	): Promise<ResponseType> => request(path, "PUT", params);
	const Delete = <RequestType, ResponseType>(
		path: string,
		params: RequestType,
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
