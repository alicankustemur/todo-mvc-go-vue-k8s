// Given : ToDo list with "rest for a while" and "drink water" item in order
// When  : I click on delete button next to "rest for a while" item
// Then  : I should see "drink water" item in ToDo list

describe('I click on delete button next to "rest for a while" item', () => {
    before('ToDo list with "rest for a while" and "drink water" item in order', () => {
        cy.server()
        cy.route('/todos',[
            {
                "id"        : 1,
                "title"     : "rest for a while",
                "order"     : 1,
                "completed" : false,
                "url"       : "http://localhost:8000/todos/1"   
            },
            {
                "id"        : 2,
                "title"     : "drink water",
                "order"     : 2,
                "completed" : false,
                "url"       : "http://localhost:8000/todos/2"   
            }
        ]);
        
        cy.server().route({
            method: 'DELETE',
            url: '/todos/1',
            status: 200,
            response: {}
        });
    })

    it('I should see "drink water" item in ToDo list', () => {
        cy.visit('/')
        cy.get('.todo-list > .todo')
            .contains('rest for a while')
            .next('.destroy')
            .click({force: true})
        
        cy.get('.todo-list > .todo')
            .first()
            .contains('drink water')
        
        cy.get('.todo-list .todo')
            .should('have.length', 1)

    })
})