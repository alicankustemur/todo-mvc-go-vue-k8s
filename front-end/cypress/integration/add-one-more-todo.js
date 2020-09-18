// Given : ToDo list with "buy some milk" item
// When  : I write "enjoy the assignment" to text box and click to add button
// Then  : I should see "enjoy the assignment" insterted to ToDo list below "buy some milk"

describe('I write "enjoy the assignment" to text box and click to add button', () => {
    before('ToDo list with "buy some milk" item', () => {
        cy.server()
        cy.route('/todos',[
            {
                "id"        : 1,
                "title"     : "buy some milk",
                "order"     : 1,
                "completed" : false,
                "url"       : "http://localhost:8000/todos/1"   
            }
        ]);
    })

    it('I should see "enjoy the assignment" insterted to ToDo list below "buy some milk"', () => {
        cy.visit('/')
        
        cy.server()
        cy.route({
            method: 'POST',
            url: '/todos',
            response:{
                "id":2,
                "title":"enjoy the assignment",
                "order":2,
                "completed":false,
                "url":"http://localhost:8000/todos/2"
            }
        })
        
        
        cy.get('.new-todo')
            .type('enjoy the assigment{enter}')
        
        cy.get('.todo-list > .todo')
            .first()
            .next()
            .contains('enjoy the assignment')
            
    })
})