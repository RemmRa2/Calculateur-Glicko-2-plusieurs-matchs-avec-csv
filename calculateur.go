package main
import(
    "math"
)
type Match struct {
    Puj float64
    RDj float64
    g float64
    E float64
    Score float64
}
type Player struct {
    Pu float64
    RD float64
    Volat float64
    Matches *[]Match 
}
func CalculateRatingFromPlayer(p Player) (Player, error) {
    
    ConvertedPu := (p.R - 1500) / 173.7178
    ConvertedRD := p.RD / 173.7178
    g := func(RD){
        return 1 / math.Sqrt(1 + 3 * (RD * RD) / math.Pi) 
    }
    E := func(Pu, Puj, RDj) {
        return 1 / (1 + math.Exp(-g(RDj) * (Pu - Puj)))
    }
    for i, m := range p.Matches {
        p.Matches[i].g = g(m.RDj) 
        p.Matches[i].E = E(m.Pu, )
    }
    NewPlayer := Player{ R: p.R, RD: p.RD, Volat: p.Volat, Matches: p.Matches }
    
}
//Do not forget give a value to tau global variable
