<script>
	import '../app.css';
</script>

<div class="drawer drawer-mobile">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content pl-4">
		<slot />
		<label for="my-drawer" class="btn btn-primary drawer-button lg:hidden">Show Sidebar</label>
	</div>
	<div class="drawer-side">
		<label for="my-drawer" class="drawer-overlay" />
		<ul class="menu p-4 w-80 bg-base-100 text-base-content">
			<li><a href="/">Home</a></li>
		</ul>
	</div>
</div>
