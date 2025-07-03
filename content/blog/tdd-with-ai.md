# TDD with AI

*Published: July 1, 2025*  
*Category: AI, TDD,*  
*Status: draft*

What’s a good workflow to write tests using AI when you simultaneously believe that the practise of TDD is useful, and you’d like to have the productivity boost of AI?

Here are three approaches I have seen to writing tests using AI and my thoughts on them. 

### Use AI to write the code and then write the tests
In this scenario the engineer probably doesn't practice TDD and in addition now also uses AI. They write the production code AI, then as the code already exists prompt the AI to write the suite of tests based on that code which can result in less focused attention on the specifics of each scenario.
Whilst a very convenient approach which requires minimal prompting I see a few dangers with this approach:

- AI can very confidently give a veneer of providing full test coverage, but be subtly incorrect.
- As the engineer didn't think through the scenarios, they have less depth of understanding of the problem space, and therefore are more likely to miss subtle bugs in the generated test suite.
- It's difficult to definitively say that the code works as it should because the test suite never started with a failing test. This was true without using AI, but now AI has been used on both sides (production and test code) becomes even more pertinent.


### TDD: use AI to write the tests and then write the code
The engineer usually practices TDD and sees it's value but wants to use AI to move more quickly. They practice TDD with AI by getting it to write the first test (or upfront whole suite of tests) and then implement the code.
The AI writes some convincing test scenarios and the production code is based off of these. The pitfalls I see of this approach are:

- It's very tempting (especially when working in a language like Go which favours table tests) to get the AI to generate the entire suite of tests outright, rather than start with one test and then move on to the next one. This means deal with the tests in bulk, rather than follow the red-green-refactor cycle for each test.
- It's more difficult to get the right prompt for the AI given the code doesn't exist yet. This can result in incorrect test suites being generated and a frustrating experience trying to correct them.
- Whilst the tests are reviewed by the engineer, it ultimately removes the thinking part involved in writing the test. My experience has been that as such I've had a shallower understanding of the code as I haven't had the creative experience of figuring out what the publicly testable interface should be.

### TDD: you write at least the first test, then prompt AI to write the rest (recommended)
Same as the above scenario, the engineer wants to both practice TDD but also use AI to move more quickly. They begin by writing the first happy path test following the red-green-refactor cycle, perhaps incorporating AI at the refactor stage and as an aide to write the production code. The engineer then makes a list of any remaining happy path tests and prompts the AI to write them based off of the initial test.
This process is then repeated for the sad path scenarios: the engineer writing the first sad path tests themselves then prompting AI to write the remaining scenarios.
Once this has been completed, they might then use code coverage tools/ AI to check if there are any missing scenarios. The challenge with this approach is that it could take more time than the first two approaches. 

The question is: why bother to write the first test as an engineer when we have AI? I've found this is an important step because:

- It maintains the value add of the TDD process: forcing the engineer to sketch out the world in the test file first and in the process have a deeper understanding of the problem space.
- It is an important creative step that will force you to answer and understand fundamental questions about the code, e.g. what is the optimal public interface? What do you really need to test?
- You are likely to more deeply interrogate the code produced by AI because you wrote the initial test and have less temptation to accept everything proposed by the AI.

### Conclusions
Overall my take on writing tests with AI is that I think it's important to write at least the first test to provide a model for the AI to follow and copy.
The reason for this is that I still want to do the groundwork thinking of what scenarios I want to test and what their intent is because this shapes the production code that gets written. 
Ultimately how an engineer uses AI with tests will depend on what they believe the value of a test is. If they currently don’t value or practice TDD, then it's likely that using AI will simply reinforce the world view that tests don't add much to the process of writing code and instead are necessary artefacts added at the end as quickly as possible.

In my case I do find it valuable to do the thinking part of writing the test first rather than getting AI to write it. I've found if I don't do that then I am less likely to interrogate the scenarios and testing approach in depth and more likely to just accept what the AI gives me.
I think this has a benefit for when the production code is written as well, because it’s clearer in my mind what I'm aiming to achieve to make sure I'm thinking critically even if using AI as a productivity aid.

I've found it's too easy to use AI to get things done more quickly at the expense of depth of understanding. However, everything has a trade-off and nothing is for free AI included. If we are trading speed for understanding then I think that speed has the potential to be dangerous for the long term.
Taking the time to understand the problem space and truly understand the code remains very important, if not more important in the age of AI.

---