# TDD with AI

*Published: July 1, 2025*  
*Category: AI, TDD,*  
*Status: draft*

So what’s a good workflow to practice TDD when you simultaneously that the practise of TDD is useful for the engineer, and at the same time you’d like to have the productivity boost of using AI.
My take on this is that I think that it's important for me to write at least the first test to provide a model for the AI to follow and copy. The reason for this is that I still want to do the groundwork thinking of what scenarios I want to test and what their intent is.

Here are three approaches I have seen to writing tests using AI:

### Use AI to write the code and then write the tests
Instead of using the test to figure out what code should be written, the engineer starts writing the production code itself and this process is made all the faster with AI. It can be very convenient to write tests in this way with AI, because once the code already exists you can just prompt the AI tool to write the tests based on that code.
In this case the engineer probably wasn’t already practicing TDD, which suggests they don't get much value from the process of deciding what to write in the tests before you implementing the code. With the advent of AI, this person can now move even more quickly. The danger that I see here is that AI tools are very confident in what they suggest. They can very confidently write all the test scenarios after writing the code and have the veneer of providing full test coverage, when in fact the test scenarios are subtly wrong.
It's also difficult to definitely say that the code works as it should because it never started with a failing test in the first place.

### Use AI to write the tests and then the code
The engineer usually practices TDD and sees it's value, but wants to use AI to move more quickly. They get AI to write the first test or the whole suite of tests and then implement the code.
The AI writes some convincing test scenarios and the production code is based off of these. The challenge that I've found with this is that it's more difficult to prompt the AI to write the test I want when the code doesn't yet exist.
Whilst the tests are reviewed by the engineer, it removes the thinking part involved in writing the test. I feel this process results in a shallower understanding of the code and the creativity involved in figuring out the publicly testable interface.

### You write at least the first test and then prompt AI follow that model
The engineer thinks about the happy path scenario and writes the first failing test. Then following the TDD cycle writes the production code to get the test to pass perhaps also using AI. The engineer then makes a list of any remaining happy path tests and prompts the AI to write them based off of the initial test.
This process is then repeated for the sad path scenarios. Why bother to do this when we have AI? I've found this is an important step because it maintains the value of the TDD process which is sketching out how the world should be in your test file. Even without the code existing, this is the important creative step that will force you to answer and understand fundamental questions about the code: how should the public interface be.
Get the test harness to look like how you want and are happy with. Give the AI to begin with one other test scenario you want it to write, get it to refer to your initial test. Work with the results until you’re happy. Now feed it the other scenarios and see that you’re happy. Be especially mindful of sad path test scenarios. It may find strange ways to mock for example. It’s worth writing this yourself first.

I've come to the conclusion that some of how an engineer will use AI with tests depends on what they believe the value of a test is. If you already don’t value TDD, it's likely that using AI will simply reinforce the world view that tests don't add much to the process of writing code and instead are necessary artefacts added at the end for the PR to be approved.
If you don’t value TDD, probably using AI will just reinforce your view of the world that they don’t add much to the process of writing code and are basically a bit pointless. 

In my case I find it valuable to force myself to do the thinking part of writing the first test and sketching out the expected sad path scenarios. Rather than getting the AI to write it and not really interrogate the scenarios and testing approach, you can be confident you fully understand the approach. This will have a benefit for when the production code is written as well, because it’s clear in your mind what you’re aiming to achieve.

It’s too easy in the age of AI to use it to get things done more quickly. However, everything has a trade off. If we are trading speed for understanding and subtle bugs then I think that speed has the potential to be dangerous.
Taking the time to understand the problem space is key and truly understand the code, even if it's partially written by AI is really important.


---