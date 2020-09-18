/*jshint unused:false */

(function (exports) {

	'use strict';
	
	const apiUrl = config.VUE_APP_API_URL;

	exports.todoStorage = {

		fetch: async function () {
			let todos = await axios.get(`${apiUrl}/todos`);
			return JSON.parse(JSON.stringify(todos.data))
		},
		save: async function (todo) {
			todo = JSON.parse(JSON.stringify(todo))

			let savedTodo = {};

			await axios.post(`${apiUrl}/todos`,{
				title: todo.title,
				completed: todo.completed
			})
			.then(function (response) {
				savedTodo = JSON.parse(JSON.stringify(response.data))
			})
			.catch(function (error) {
				console.log(error);
			});

			return savedTodo;
		},
		delete: async function (id) {
			await axios.delete(`${apiUrl}/todos/${id}`).then(response => response.data);
		},
		update: async function (id,editedTodo) {
			await axios.patch(`${apiUrl}/todos/${id}`, editedTodo )
					.then(response => response.data);
		}
	};

})(window);
