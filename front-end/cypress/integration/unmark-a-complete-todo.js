// Given   : ToDo list with marked "buy some milk" item
// When    : I click on "buy some milk" text
// Then    : I should see mark of "buy some milk" item should be cleared as "buy some milk"

describe('I click on "buy some milk" text', () => {
    before('ToDo list with "buy some milk" item', () => {
        cy.server()
        cy.route('/todos',[
            {
                "id"        : 1,
                "title"     : "buy some milk",
                "order"     : 1,
                "completed" : true,
                "url"       : "http://localhost:8000/todos/1"   
            }
        ]);
        
        cy.server().route({
            method: 'PATCH',
            url: '/todos/1',
            status: 200,
            request:{
                "title"     : "buy some milk",
                "completed" : false,
            },
            response:{
                "id":1,
                "title":"buy some milk",
                "order":1,
                "completed":false,
                "url":"http://localhost:8000/todos/1"
            }
        });
    })

    it('I should see "buy some milk" item marked as "buy some milk"', () => {
        cy.visit('/')
        cy.get('.todo-list > .todo')
            .first()
            .find('[type="checkbox"].toggle')
            .uncheck()
            .should('not.be.checked')
    })
})