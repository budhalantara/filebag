<script lang="ts">
	import task from '$lib/services/task';
	import { FileEarmarkPlus } from 'svelte-bootstrap-icons';
	import { createEventDispatcher } from 'svelte';

	let checked = false;
	let url = '';

	const dispatch = createEventDispatcher();

	async function submit() {
		await task.create({
			url,
			raw_url: url,
			connection_count: 2
		});
		checked = false;
		dispatch('added');
	}
</script>

<!-- The button to open modal -->
<label for="add-task-modal" class="btn btn-primary btn-sm mb-4">
	<FileEarmarkPlus class="mr-1" />
	New
</label>

<!-- Put this part before </body> tag -->
<input type="checkbox" id="add-task-modal" class="modal-toggle" bind:checked />
<div class="modal">
	<div class="modal-box relative">
		<label for="add-task-modal" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
		<h3 class="text-lg font-bold">New Task</h3>
		<div class="py-2">
			<div class="form-control w-full">
				<label class="label" for="url">
					<span class="label-text">URL</span>
				</label>
				<input
					type="text"
					id="url"
					placeholder="https://dummyimage.com/600x400/000/fff"
					class="input input-bordered w-full"
					bind:value={url}
				/>
			</div>
			<div class="modal-action">
				<button class="btn btn-sm" on:click={submit}>Add</button>
			</div>
		</div>
	</div>
</div>
