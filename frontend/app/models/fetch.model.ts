export interface HttpClientOption {
  url: string;
  method: string;
  data?: any; // Request body - can be JSON object, FormData, Blob, File, URLSearchParams, or string
  body?: any; // Alias for data (for fetch API compatibility)
  headers?: Record<string, string>;
  params?: Record<string, any>; // Query parameters
  responseType?: "json" | "text" | "blob" | "download" | "arrayBuffer";
}

// Keep the old interface name for backward compatibility
export interface option extends HttpClientOption {}
