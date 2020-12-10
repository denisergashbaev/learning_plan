# Reinforcement Learning in Healthcare: Challenges and Promise

* 0:42 find the right treatment for the right person
* 0:56 talk focuses on HIV. Over time patients develop resistance to the drugs and you need to switch the drugs over time. This requires that the doctors are careful when proposing the treatments. The schedule should be such that when a drug B ceases to become effective, the subsequent (not-yet tried) drugs should be thought of ahead of time
* The idea is to combine model-based RL with neighborhood-based prediction to optimize the treatment plan. Why? The models we have at hand are too simplistic and cannot represent our system that well. Therefore we rely on neighborhood-based recommendation -- ie chose a treatment plan that was successful for patients with similar profile (results, previous history, etc). Sometimes, though, we do not have good enough neighbor data and therefore need the model-based approach
  
  * 11:11 reward signal is just some model coming from collaborators. That is, the model is not the model of the environment in the traditional RL sense (prediction of how the world changes given the action and the ensuing reward). The model is a mathematical equation that determines our reward
  * We look for other patients that match our profile and see what action (drug)was taken in their case at a (similar) time point. We select the same action that was given to those patients that have been doing well 21 days into the future after the treatment was given. The neighborhood-based approach only looks into short-term. Why?
    * One action alone does not define a patient's success in the advanced future. There are many other different modalities (other treatments) that play a role in the long-term success
    * However, we use it as a starting point, by assuming that if we select the same treatment for our patient as the treatment for another patient that was doing well 21 days into the treatment, then we are off to a good start.
  * 15:34 - Combined model trained end-to-end using automatic differentiation  
    ![Combined model](images/combined-model.png?raw=true)
    * Kernel -- neighborhood-based approach (does the viral go down immediately or stays low over a longer period of time)
    * POMDP -- our model's action preference based on limited "textbook" knowledge
    * Patient statistics is used to determine which action to pick -- Kernel or POMDP
  * 16:00 Does this model work? Policy-mixture policy performs the best, preferring POMDP around 20-30% of the time.
    * Actions: 312 common drug combinations
    * RL algorithm: **not specified** (TODO: look up)
    * 17:39 POMDP is chose when histories are long and the neighbor distances are big, which matches the intuition
  * 18:30 Current work: Predicting results of the actions
  * 21:23 Applied to Sepsis management in the ICU
    * Initially, the problem looks the similar to the HIV treatment. Not quite, in HIV we want to reduce the integral of the viral load, in sepsis management we focus on whether the patient survives at the end (not that much whether their blood pressure was at a certain level)
    * Focusing on mortality as the outcome of the treatment makes it much harder to evaluate
  * 26:40 We evaluate on the test dataset by finding a group of patients whose treatment matches the predicted treatment, that makes your [test] dataset smaller
  * 30:58 think of AI as complimentary assistance for doctors (not a replacement)
  * Perfection is the enemy of good  
Questions:
17:39 When Kernel vs POMDP are chosen: Distance to 2nd quantile to your neighbors
20:58 Interpolating the model allows us to get better results
