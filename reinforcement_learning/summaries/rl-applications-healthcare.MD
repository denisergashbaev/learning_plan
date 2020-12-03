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
Finding a good reward function in RL can be a significant challenge in some environments. This issue becomes 