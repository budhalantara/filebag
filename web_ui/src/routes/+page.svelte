<script lang="ts">
	import task from '$lib/services/task';
	import type { Task } from '$lib/types/task';
	import { FileEarmarkPlus } from 'svelte-bootstrap-icons';

	let tasks: Task[] = [];

	async function load() {
		tasks = await task.getAll();
	}

	load();
</script>

<div class="py-4">
	<h1 class="font-bold text-2xl">Home</h1>
	<h2>Download List</h2>
</div>

<button class="btn btn-primary btn-sm mb-4">
	<FileEarmarkPlus class="mr-1" />
	New
</button>

<div class="overflow-x-auto">
	<table class="table table-compact w-full">
		<thead>
			<tr>
				<th>#</th>
				<th>File Name</th>
				<th>File Size</th>
				<th>Status</th>
				<th>Date</th>
			</tr>
		</thead>
		<tbody>
			{#each tasks as task, i}
				<tr>
					<th>{i + 1}</th>
					<td>{task.file_name}</td>
					<td>{task.file_size}</td>
					<td>{task.status}</td>
					<td>{task.date}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
