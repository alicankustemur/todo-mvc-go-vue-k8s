// Given : ToDo list with "rest for a while" item
// When  : I click on delete button next to "rest for a while" item
// Then  : List should be empty

describe('I click on delete button next to "rest for a while" item', () => {
    before('ToDo list with "rest for a while" item', () => {
        cy.server()
        cy.route('/todos',[
            {
                "id"        : 1,
                "title"     : "rest for a while",
                "order"     : 1,
                "completed" : false,
                "url"       : "http://localhost:8000/todos/1"   
            }
        ]);
        
        cy.server().route({
            method: 'DELETE',
            url: '/todos/1',
            status: 200,
            response: {}
        });
    })

    it('List should be empty', () => {
        cy.visit('/')
        cy.get('.todo-list > .todo')
            .contains('rest for a while')
            .next('.destroy')
            .click({force: true})

        cy.get('.todo-list .todo')
            .should('have.length', 0)
    })
})
