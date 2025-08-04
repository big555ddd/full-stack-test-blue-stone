import type { option } from "~/models/fetch.model";

export const client = async (ops: option): Promise<any> => {
  const config: any = useRuntimeConfig();

  const baseURL =
    config.public.WEB_API?.replace(/\/$/, "") + "/api/v1" ||
    "http://localhost:8080/api/v1";

  // Construct full URL
  let url = `${baseURL}${ops.url}`;

  // Handle query parameters
  if (ops.params) {
    const searchParams = new URLSearchParams();
    Object.entries(ops.params).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        searchParams.append(key, String(value));
      }
    });
    const queryString = searchParams.toString();
    if (queryString) {
      url += `?${queryString}`;
    }
  }

  // Prepare headers
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...ops.headers,
  };

  // Prepare request options
  const requestOptions: RequestInit = {
    method: ops.method.toUpperCase(),
    headers,
  };

  // Use body or data (body takes precedence for fetch API compatibility)
  const requestBody = ops.body || ops.data;

  // Handle request body
  if (requestBody && !["GET", "HEAD"].includes(ops.method.toUpperCase())) {
    // Handle FormData (for file uploads)
    if (requestBody instanceof FormData) {
      delete headers["Content-Type"]; // Let browser set the boundary for FormData
      requestOptions.body = requestBody;
    }
    // Handle Blob or File
    else if (requestBody instanceof Blob || requestBody instanceof File) {
      requestOptions.body = requestBody;
    }
    // Handle URLSearchParams (for form-encoded data)
    else if (requestBody instanceof URLSearchParams) {
      headers["Content-Type"] = "application/x-www-form-urlencoded";
      requestOptions.body = requestBody.toString();
    }
    // Handle string data
    else if (typeof requestBody === "string") {
      requestOptions.body = requestBody;
    }
    // Handle JSON data (default)
    else if (headers["Content-Type"] === "application/json") {
      requestOptions.body = JSON.stringify(requestBody);
    }
    // Handle other data types
    else {
      requestOptions.body = requestBody;
    }
  }

  try {
    const response = await fetch(url, requestOptions);

    // Handle download/blob responses
    if (ops.responseType === "download" || ops.responseType === "blob") {
      const blob = await response.blob();
      return {
        data: blob,
        status: response.status,
        statusText: response.statusText,
        headers: response.headers,
      };
    }

    // Handle JSON responses
    let data;
    try {
      data = await response.json();
    } catch {
      data = await response.text();
    }

    // Return axios-like response format for compatibility
    return {
      data,
      status: response.status,
      statusText: response.statusText,
      headers: response.headers,
      config: ops,
    };
  } catch (error) {
    throw {
      message: error instanceof Error ? error.message : "Network error",
      status: 0,
      statusText: "Network Error",
      config: ops,
    };
  }
};
