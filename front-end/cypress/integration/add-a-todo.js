// Given : Empty ToDo list
// When  : I write "buy some milk" to text box and click to add button
// Then  : I should see "buy some milk" item in ToDo list

describe('I write "buy some milk" to text box and click to add button', () => {
    
    before('Empty ToDo list', () => {
        cy.server().route({
            method: 'GET',
            url: '/todos',
            status: 200,
            response: [
            ],
        });
        
        cy.server().route({
            method: 'POST',
            url: '/todos',
            status: 200,
            response:{
                "id":1,
                "title":"buy some milk",
                "order":1,
                "completed":false,
                "url": "http://localhost:8000/todos/1"
            }
        });
    })

    it('I should see "buy some milk" item in ToDo list', () => {
        cy.visit('/')

        cy.get('.new-todo')
            .type('buy some milk{enter}')

        cy.get('.todo-list > .todo')
            .find('label')
            .contains('buy some milk')
        
        cy.get('.todo-list .todo')
            .should('have.length', 1)
    })
})