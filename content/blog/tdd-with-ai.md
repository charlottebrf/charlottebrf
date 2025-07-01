# TDD with AI

*Published: July 1, 2025*  
*Category: AI, TDD,*  
*Status: draft*

So what’s a good workflow to do TDD when you simultaneously believe that software engineers have value and that you’d like to have a productivity value add of using AI. I think it is important for the engineer to write the first test which provides a model for the to follow and copy. Writing a TDD test then I would still want to do the groundwork of considering:

- What scenarios to test
- Writing the first test myself
- Getting the AI to write the scenarios you’ve written
- I craft the intent of the tests

Some of how you use AI with TDD depends on what you believe the value of a test is in a codebase. My theory is that using AI tools makes doing TDD more difficult or seemingly pointless, especially if the engineer wasn’t already practising TDD.  All of these approaches assume that the engineer will be working with

Some examples of what I’ve seen in the wild:

- Working in haste. Whilst figuring out what code to write, they actually dive in and write the code. They then write the tests, and given the code already exists they use an AI tool to write the tests. In this case the engineer already wasn’t doing TDD, so doesn’t value the purpose of TDD as a core thinking part of the process of writing code. With the advent of AI, this person can now move even more quickly. Writing the tests are a bit of an arbitrary exercise. They need to be added in order for the PR to be approved, and an after-thought of the code.
- Write the tests first, then write the code all using AI.  The challenge with this approach is that the tool doesn’t have ready written code as a source for it to use to write the tests.  How do you prompt it to think about the scenarios that you need to test. How do you get it to understand the intent. Think about the happy path scenario and then sad path scenarios you want to test. Sketch out how the world should be in your test file first. Even without the code existing, this is the important and imaginative step where humans can excel. It will force you to answer questions: how should the interface be? What kind of test are you writing and therefore what makes sense to include in the scope of this test? What are the sad path tests? Seriously: you do the groundwork thinking to consider these , again with the considerations of which test you’re writing. Write the first test yourself without the assistance of the AI. Go through a single TDD cycle, get the test passing. Ask the AI to make suggestions for refactoring. Get the test harness to look like how you want and are happy with. Give the AI to begin with one other test scenario you want it to write, get it to refer to your initial test. Work with the results until you’re happy. Now feed it the other scenarios and see that you’re happy. Be especially mindful of sad path test scenarios. It may find strange ways to mock for example. It’s worth writing this yourself first. Even if you go with AI suggestions, thinking through how & why it’s chosen an approach is a worthwhile exercise.

Even if you end up with code that’s exactly the same as what an AI would have written, this process is still valuable. Rather than getting the AI to write it and not really interrogate the scenarios and testing approach, you can be confident you fully understand the approach. This will have a benefit for when the production code is written as well, because it’s clear in your mind what you’re aiming to achieve.

If you don’t value TDD, probably using AI will just reinforce your view of the world that they don’t add much to the process of writing code and are basically a bit pointless. Whilst reviewing PRs you’ll just continue to quickly gloss over the tests.

It’s too easy in the age of AI to use it to get things done more quickly. However, in my experience, this has often been a fallacy. Taking the time to understand the problem space is key. That way you can be sure that you can interrogate the AI and call BS when you see it.

---