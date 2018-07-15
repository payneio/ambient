package reflex

// TODO: Support the following kinds of rules:
// - anyone is in the workshop => tv:on, overhead light: on, workbench light: on and workstations:awake
// - nobody is in the workship => tv:off, overhead light: off, workbench light: off and workstations:asleep

// maintain immutable state trees: actual, desired
// when actual state matches X, issue command
// commands transform actual state to a new desired state
// ambientd then takes care of achieving desired state
