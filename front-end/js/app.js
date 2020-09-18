/*global Vue, todoStorage */

(function (exports) {
	'use strict';
	
	var filters = {
		all: function (todos) {
			return todos;
		},
		active: function (todos) {
			return todos.filter(function (todo) {
				return !todo.completed;
			});
		},
		completed: function (todos) {
			return todos.filter(function (todo) {
				return todo.completed;
			});
		}
	};

	exports.app = new Vue({


		// the root element that will be compiled
		el: '.todoapp',

		// app initial state
		data: {
			todos: [],
			newTodo: '',
			editedTodo: null,
			visibility: 'all'
		},

		mounted: async function(){
			this.todos = await todoStorage.fetch()
		},

		// computed properties
		// http://vuejs.org/guide/computed.html
		computed: {
			filteredTodos: function () {
				return filters[this.visibility](this.todos);
			},
			remaining: function (value) {
				return filters.active(this.todos).length;
			},
			allDone: {
				get: function () {
					return this.remaining === 0;
				},
				set: function (value) {
				let	changeCompleted = this.changeCompleted;

					this.todos.forEach(function (todo) {
						console.log(todo);
						todo.completed = value;
						changeCompleted(todo);
					});
				}
			}
		},

		// methods that implement data logic.
		// note there's no DOM manipulation here at all.
		methods: {

			pluralize: function (word, count) {
				return word + (count === 1 ? '' : 's');
			},

			addTodo: async function () {
				var value = this.newTodo && this.newTodo.trim();
				if (!value) {
					return;
				}
				let todo = { title: value, completed: false };
				let savedTodo = await todoStorage.save(todo);
				this.todos.push(savedTodo);

				this.newTodo = '';
			},

			removeTodo: async function (todo) {
				await todoStorage.delete(todo.id);
				var index = this.todos.indexOf(todo);
				this.todos.splice(index, 1);
			},

			editTodo: function (todo) {
				this.beforeEditCache = todo.title;
				this.editedTodo = todo;
			},

			changeCompleted: async function (todo) {
				if (!todo) {
					return;
				}
				await todoStorage.update(todo.id, todo);
			},

			doneEdit: async function (todo) {
				if (!this.editedTodo) {
					return;
				}
				await todoStorage.update(todo.id, this.editedTodo);

				this.editedTodo = null;
				todo.title = todo.title.trim();
				if (!todo.title) {
					this.removeTodo(todo);
				}
			},

			cancelEdit: function (todo) {
				this.editedTodo = null;
				todo.title = this.beforeEditCache;
			},

			removeCompleted: function () {
				let allTodos = filters.all(this.todos);
				let completedTodos = filters.completed(this.todos);

				let	removeTodo = this.removeTodo;
				completedTodos.forEach(function (todo) {
					removeTodo(todo);
				});

				this.todos = allTodos;

			},

		},

		// a custom directive to wait for the DOM to be updated
		// before focusing on the input field.
		// http://vuejs.org/guide/custom-directive.html
		directives: {
			'todo-focus': function (el, binding) {
				if (binding.value) {
					el.focus();
				}
			}
		}
	});

})(window);
