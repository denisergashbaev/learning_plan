# A review of recent reinforcement learning applications to healthcare

In healthcare, machine learning is typically used for diagnosis or outcome prediction, while optimizing treatment receives less attention. In this respect, reinforcement learning (RL) is a promising technique as it focuses on the optimization of the policy (ie, taking decisions according to given states that bring about the best cumulative outcome). However, RL techniques have achieved notable results when applied to games, where the problems of data availability, observability, reward sparsity are oftentimes of lesser intensity than in real-world situations.

Some of the challenges of applying RL in healthcare settings are addressed in this section:

* **Learning and evaluating policy on the historical observational data**  
Unlike as it is the case with games where RL agent acts and changes the environment while learning, it is usually not possible to have the RL system learn from live data in the medical environment. Neither is evaluating of the algorithm's performance on live data is possible. Although Q-Learning can be used to learn the optimal policy in the off-policy context, evaluating the learned policy is tricky.  
Several evaluation metrics are discussed:
  
  * Importance sampling. **TODO**: describe importance sampling here
  * U-Curve

* **Partial observability**  
Compared with games, the issue of partial observability in healthcare from the RL perspective is arguably more severe. There is only a very limited set of measurements coming from the patient -- eg, blood pressure, temperature, SO2, plus some simple measures -- that do not give enough information about the person. In addition, RL system may need to deal with extremely sparse measurements, such as X-rays taken before and after treatment of pneumonia.

* **Reward function**  
Finding a good reward function in RL can be a significant challenge in some environments. This issue becomes more prominent in the healthcare setting: 

  * Hard to find non-sparse reward function (e.g. short term improvement in healthcare may not mean improvement in outcome in case of sepsis). Reward in the case of death/survival is too sparse. (Denis: the may be different states of survival, how would one model those with a reward function)

* **RL is data hungry (i.e., sample efficiency)**

All major breakthroughs were achieved in settings where the amounts of data used would correspond to years of human experience. Medical data (for instance, treatments) is very scarce.

* **Non-stationarity of data**
Healthcare data is non-stationary and dynamic:

  * Symptoms recorded at inconsistent intervals (e.g., immediately after the treatment or with delay, different frequency)
  * Changing treatment objectives depending on the progress of disease and accompanying diseases

## Interesting recent studies

### Deep Reinforcement Learning for Sepsis Treatment

* Uses sepsis subset of the MIMIC-III dataset
* Action space: Vasopressors and IV fluid
  * drug doses grouped in four bins
* Double DQN with separate value and advantage heads
* Reward function based on SOFA score
* Evaluation: U-Curve --- mortality rate as a function between prescribed policy vs actual policy (TODO: look it up, what is prescribed and actual)

### Reinforcement Learning with Action-Derived Rewards for Chemotherapy and Clinical Trial Dosing Regimen Selection

* Objective: find optimal policy for chemotherapy
* Q-Learning (TODO: DQN?)
* Action space: quantity of doses for given duration

  * dose cycles initiated with frequency determined by experts
  * transitions computed at end of each cycles

* Reward function: mean reduction in tumour diameter
* Evaluation: simulated clinical trials

### Supervised Reinforcement Learning with Recurrent Neural Network for Dynamic Treatment Recommendation

* Supervised RL (via Actor Critic methods) with RNN:

  * Actor recommends medications
  * Critic estimates the value of medications
  * LSTM alleviates partial observability by summarizing the historical observations

* Uses full MIMIC-III dataset
* Action space: 100 exact medications or 180 drug categories
* Evaluation: based on estimated in-hospital mortality (TODO)

### Deep Reinforcement Learning for Dynamic Treatment Regimes on Medical Registry Data

* Proposes a framework:

  * supervised learning to predict most possible expert treatments
  * RL minimizes possibility of complications for the treatment

### Reinforcement Learning System to Encourage Physical Activity in Diabetes Patients

* TODO : read & summarize
* encourages healthy habits instead of treatments

### Representation and Reinforcement Learning for Personalized Glycemic Control in Septic Patients

### Learning Optimal Policies form Observational Data

* Not RL but uses counterfactual inference and domain adversarial NN