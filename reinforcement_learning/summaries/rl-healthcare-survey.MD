# RL in healthcare survey

* RL does not require a mathematical model; contrast this to control-based approaches in healthcare

1. Efficient techniques
    * BatchRL is an *experience-level* approach. It uses two techniques: 1) storing experience in a buffer and reusing it as if it were new, 2) separating DP (dynamic programming) step from the function approximation step by using a supervised learning to fit the function approximator over the sampled experience. Famous approaches:
        * Neural Fitted Q Iteration
        * Tree-based FQI
        * Least-Squares Policy Iteration
    These methods have been used in the medical setting

    * *Model-based RL* methods build a model of the environment: that is, they learn transition probability and reward function. (**TODO: how they construct the model**) This model is used with advanced exploration strategies (**TODO: what kind?**) and planning (Monte-Carlo Tree Search (MCTS))

    * *Task-level*, *Transfer RL* uses transfer knowledge to improve the learning performance in the target task (by reducing the number of samples or increasing hte final convergence level)

2. Representational techniques
    * *Representation for Value Functions or Policies*
    * *Representation for Reward Functions*
    * *Representation for Tasks or Models*

## RL in Healthcare

Applied in a number of healthcare domains:

* Dynamic Treatment Regimes (alternatively named, dynamic treatment policies, adaptive interventions, adaptive treatment strategies)
  * Chronic Diseases
  * Critical Care
* Automated Medical Diagnosis
  * Structured Data
  * Unstructured Data
* Other General Domains
  * Resource Scheduling and Allocation
  * Optimal Process Control
  * Drug Discovery
  * Health Management