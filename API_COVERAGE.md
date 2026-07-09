# CHPP file coverage

This table was last fully audited on 2026-07-09.

**Status legend**
- ✅ Yes — audited against the doc, correct and complete, wired into `api`/`Client`
- 🐛 Bug — has real content and is wired, but has one or more confirmed correctness bugs (silently wrong/missing data, wrong types, unwired actionTypes/params, or a stale version). See the plan for specifics.
- ⚠️ Stub — `chpp/file_*.go` exists but only has the generic envelope (`FileName`/`Version`/`User`/`Error`/...), no real payload, and no `api`/`Client` wiring
- ❌ No — no file at all in this repo

| CHPP file | Latest API version | Status | Current version in repo |
|---|---|---|---|
| achievements | 1.2 | ✅ Yes | 1.2 |
| alliancedetails | 1.5 | ✅ Yes | 1.5 |
| alliances | 1.4 | ✅ Yes | 1.4 |
| arenadetails | 1.7 | ✅ Yes | 1.7 |
| avatars | 1.1 | ✅ Yes | 1.1 |
| bookmarks | 1.0 | ✅ Yes | 1.0 |
| challenges | 1.6 | ✅ Yes | 1.6 |
| club | 1.5 | ✅ Yes | 1.5 |
| cupmatches | 1.4 | ✅ Yes | 1.4 |
| currentbids | 1.0 | ✅ Yes | 1.0 |
| economy | 1.4 | ✅ Yes | 1.4 |
| fans | 1.3 | ✅ Yes | 1.3 |
| hofplayers | 1.2 | ✅ Yes | 1.2 |
| ladderdetails | 1.0 | ✅ Yes | 1.0 |
| ladderlist | 1.0 | ✅ Yes | 1.0 |
| leaguedetails | 1.6 | ✅ Yes | 1.6 |
| leaguefixtures | 1.2 | ✅ Yes | 1.2 |
| leaguelevels | 1.0 | ✅ Yes | 1.0 |
| live | 2.3 | ✅ Yes | 2.3 |
| managercompendium | 1.7 | ✅ Yes | 1.7 |
| matches | 2.9 | ✅ Yes | 2.9 |
| matchesarchive | 1.5 | ✅ Yes | 1.5 |
| matchdetails | 3.1 | ✅ Yes | 3.1 |
| matchlineup | 2.1 | ✅ Yes | 2.1 |
| matchorders | 3.1 | ✅ Yes | 3.1 |
| nationalplayers | 1.5 | ✅ Yes | 1.5 |
| nationalteamdetails | 1.9 | ✅ Yes | 1.9 |
| nationalteammatches | 1.4 | ✅ Yes | 1.4 |
| nationalteams | 1.6 | ✅ Yes | 1.6 |
| players | 2.8 | ✅ Yes | 2.8 |
| playerdetails | 3.2 | ✅ Yes | 3.2 |
| playerevents | 1.3 | ✅ Yes | 1.3 |
| regiondetails | 1.2 | ✅ Yes | 1.2 |
| search | 1.2 | ✅ Yes | 1.2 |
| staffavatars | 1.1 | ✅ Yes | 1.1 |
| stafflist | 1.2 | ✅ Yes | 1.2 |
| supporters | 1.0 | ✅ Yes | 1.0 |
| teamdetails | 3.9 | ✅ Yes | 3.9 |
| tournamentdetails | 1.0 | ✅ Yes | 1.0 |
| tournamentfixtures | 1.1 | ✅ Yes | 1.1 |
| tournamentleaguetables | 1.1 | ✅ Yes | 1.1 |
| tournamentlist | 1.0 | ✅ Yes | 1.0 |
| training | 2.2 | ✅ Yes | 2.2 |
| trainingevents | 1.3 | ✅ Yes | 1.3 |
| transfersearch | 1.1 | ✅ Yes | 1.1 |
| transfersplayer | 1.1 | ✅ Yes | 1.1 |
| transfersteam | 1.2 | ✅ Yes | 1.2 |
| translations | 1.2 | ✅ Yes | 1.2 |
| worldcup | 1.1 | ✅ Yes | 1.1 |
| worlddetails | 2.0 | ✅ Yes | 2.0 |
| worldlanguages | 1.2 | ✅ Yes | 1.2 |
| youthavatars | 1.2 | ✅ Yes | 1.2 |
| youthleaguedetails | 1.1 | ✅ Yes | 1.1 |
| youthleaguefixtures | 1.0 | ✅ Yes | 1.0 |
| youthplayerdetails | 1.3 | ✅ Yes | 1.3 |
| youthplayerlist | 1.3 | ✅ Yes | 1.3 |
| youthteamdetails | 1.4 | ✅ Yes | 1.4 |

**Summary:** 57 clean, 0 implemented-but-buggy, 0 stubs, 0 missing. 57 CHPP files total — full coverage.
