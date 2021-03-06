package parser

import(
	"testing"
	"github.com/jptosso/coraza-waf/pkg/engine"
)

func TestString(t *testing.T) {
	rule := `SecRule ARGS:/(.*?)/|REQUEST_HEADERS|!REQUEST_HEADERS:/X-(Coraza|\w+)/ "@rx (.*?)" "id:1, drop, phase: 1"`
	waf := &engine.Waf{}
	waf.Init()
	p := &Parser{}
	p.Init(waf)
	p.Evaluate(rule)
	
	if len(waf.Rules.GetRules()) != 1{
		t.Error("Rule not created")
	}
	r := waf.Rules.GetRules()[0]
	if len(r.Actions) != 3{
		t.Error("Failed to parse actions")
	}
	if len(r.Variables) != 2{
		t.Error("Failed to parse variables, got", len(r.Variables))
	}
	if len(r.Variables[1].Exceptions) != 1{
		t.Error("Failed to add exceptions to rule variable")
		return
	}
	if r.Variables[1].Exceptions[0] != `/x-(coraza|\w+)/`{
		t.Error("Invalid variable key for regex, got:", r.Variables[1].Exceptions[0])
	}
}

func TestString2(t *testing.T){
	rule := `SecRule ARGS|ARGS_NAMES|REQUEST_COOKIES|!REQUEST_COOKIES:/__utm/|REQUEST_COOKIES_NAMES|REQUEST_BODY|REQUEST_HEADERS|XML:/*|XML://@* \
    "@rx (?:rO0ABQ|KztAAU|Cs7QAF)" \
    "id:944210,\
    phase:2,\
    block,\
    log,\
    msg:'Magic bytes Detected Base64 Encoded, probable java serialization in use',\
    logdata:'Matched Data: %{MATCHED_VAR} found within %{MATCHED_VAR_NAME}',\
    tag:'application-multi',\
    tag:'language-java',\
    tag:'platform-multi',\
    tag:'attack-rce',\
    tag:'OWASP_CRS',\
    tag:'OWASP_CRS/WEB_ATTACK/JAVA_INJECTION',\
    tag:'WASCTC/WASC-31',\
    tag:'OWASP_TOP_10/A1',\
    tag:'PCI/6.5.2',\
    tag:'paranoia-level/2',\
    ver:'OWASP_CRS/3.2.0',\
    severity:'CRITICAL',\
    setvar:'tx.rce_score=+%{tx.critical_anomaly_score}',\
    setvar:'tx.anomaly_score_pl2=+%{tx.critical_anomaly_score}'"`

	waf := &engine.Waf{}
	waf.Init()
	p := &Parser{}
	p.Init(waf)
	p.Evaluate(rule)
	
	if len(waf.Rules.GetRules()) != 1{
		t.Error("Rule not created")
		return
	}
	r := waf.Rules.GetRules()[0]
	if len(r.Variables) != 8{
		t.Error("Failed to parse variables, got", len(r.Variables))
		for _, v := range r.Variables{
			t.Error(v)
		}
	}
}

func TestString3(t *testing.T) {
	rule := `SecRule REQUEST_HEADERS:User-Agent "@rx (.*?)" "id:1, drop, phase: 1"`
	waf := &engine.Waf{}
	waf.Init()
	p := &Parser{}
	p.Init(waf)
	p.Evaluate(rule)
	
	if len(waf.Rules.GetRules()) != 1{
		t.Error("Rule not created")
	}
	r := waf.Rules.GetRules()[0]
	if len(r.Actions) != 3{
		t.Error("Failed to parse actions")
	}
	if len(r.Variables) != 1 && r.Variables[0].Key != "User-Agent"{
		t.Error("Failed to parse variables")
	}
}

/*
* Directives
* TODO There should be an elegant way to separate them from the parser
*/

