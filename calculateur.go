package main
type Match struct {
    puj float64
    RDj float64
    g float64
    E float64
    Score float64
}
type Player struct {
    R float64
    RD float64
    Volat float64
    Matches *[]Match 
}
func CalculateRatingFromPlayer(p Player) (Player, error) {
    newPlayer := Player{ R: p.R, RD: p.RD, Volat: p.Volat, Matches: p.Matches }
    pu := (p.R - 1500) / 173.7178
    RD := p.RD / 173.7178
    
}
//Do not forget give a value to tau global variable
