import axios from '$lib/axios';
import type { CreateTaskParams, Task, TaskResponse } from '$lib/types/task';
import dayjs from 'dayjs';
import byteSize from 'byte-size';

async function getAll(): Promise<Task[]> {
	const res = await axios.get('/api/tasks');
	return res.data?.data?.map((task: TaskResponse) => {
		const data: Task = {
			...task,
			file_size: byteSize(task.file_size).toString(),
			status: task.status.substring(0, 1).toUpperCase() + task.status.substring(1),
			date: dayjs.unix(task.created_at).format('DD MMMM YYYY HH:mm:ss')
		};
		return data;
	});
}

async function create(params: CreateTaskParams) {
	await axios.post('/api/tasks', params);
}

export default { getAll, create };
