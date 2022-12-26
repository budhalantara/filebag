export type TaskResponse = {
	id: number;
	url: string;
	raw_url: string;
	file_name: string;
	file_size: number;
	connection_count: number;
	status: string;
	created_at: number;
};

export type Task = {
	id: number;
	url: string;
	raw_url: string;
	file_name: string;
	file_size: string;
	connection_count: number;
	status: string;
	date: string;
};
