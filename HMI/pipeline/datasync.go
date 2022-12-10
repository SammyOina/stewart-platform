package pipeline

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/gorilla/websocket"
	"github.com/sammyoina/stewart-platform-ui/api"
	"github.com/sammyoina/stewart-platform-ui/calibration"
	"github.com/sammyoina/stewart-platform-ui/fileWriter"
	"github.com/sammyoina/stewart-platform-ui/models"
	"github.com/sammyoina/stewart-platform-ui/queue"
	"github.com/sammyoina/stewart-platform-ui/state"
	"google.golang.org/protobuf/proto"
)

var (
	Fx                  []float64
	Fy                  []float64
	Fz                  []float64
	Mx                  []float64
	My                  []float64
	Mz                  []float64
	IntakeVelocity      []float64
	TestSectionVelocity []float64
	DiffuserVelocity    []float64
	Strain1             []float64
	Strain2             []float64
	Strain3             []float64
	Strain4             []float64
	Strain5             []float64
	Strain6             []float64
	Yaw                 []float64
	Pitch               []float64
	Roll                []float64
	RecordData          bool = false
	IMUWriter           *fileWriter.FileWriter
	PitotWriter         *fileWriter.FileWriter
	StrainWriter        *fileWriter.FileWriter
	ForceMomentsWriter  *fileWriter.FileWriter
)

type STDSync struct {
}

var currentState state.SystemState

func init() {
	currentState = *state.NewState()
}

func (o *STDSync) StartOutputting(q queue.Queue) {
	fmt.Println("Starting output")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			var e models.SensorEvent
			if err := proto.Unmarshal(message, &e); err != nil {
				log.Println("failed to unmarshal:", err)
				return
			}
			switch event := e.Event.(type) {
			case *models.SensorEvent_IMUEvent:
				if RecordData && IMUWriter != nil {
					RecordToFile(IMUWriter, event.IMUEvent.Yaw, event.IMUEvent.Pitch, event.IMUEvent.Roll)
				}
				imuToUi(event.IMUEvent.Yaw, event.IMUEvent.Pitch, event.IMUEvent.Roll)
				//fmt.Println("got data: ", event.IMUEvent.Pitch, event.IMUEvent.Yaw, event.IMUEvent.Roll)
			case *models.SensorEvent_PitotEvent:
				if RecordData && PitotWriter != nil {
					RecordToFile(PitotWriter, event.PitotEvent.DiffuserPitot, event.PitotEvent.IntakePitot, event.PitotEvent.TestSectionPitot)
				}
				pitotToUi(event.PitotEvent.DiffuserPitot, event.PitotEvent.IntakePitot, event.PitotEvent.TestSectionPitot)
				//fmt.Println("got data: ", event.PitotEvent.DiffuserPitot, event.PitotEvent.IntakePitot, event.PitotEvent.TestSectionPitot)
			case *models.SensorEvent_StrainEvent:
				if RecordData && StrainWriter != nil {
					RecordToFile(StrainWriter, event.StrainEvent.Strain1, event.StrainEvent.Strain2, event.StrainEvent.Strain3, event.StrainEvent.Strain4, event.StrainEvent.Strain5, event.StrainEvent.Strain6)
				}
				forcesToUi(event.StrainEvent.Strain1, event.StrainEvent.Strain2, event.StrainEvent.Strain3, event.StrainEvent.Strain4, event.StrainEvent.Strain5, event.StrainEvent.Strain6)
				//fmt.Println("got data: ", event.StrainEvent.Strain1, event.StrainEvent.Strain2, event.StrainEvent.Strain3, event.StrainEvent.Strain4, event.StrainEvent.Strain5, event.StrainEvent.Strain6)
			default:
				fmt.Println("no sensor event found")
				fmt.Println(hex.EncodeToString(message))
			}
		}
	}
}

func pitotToUi(diffuser float32, intake float32, testSection float32) {
	var k int = 1
	IntakeVelocity = append(IntakeVelocity[k:], IntakeVelocity[0:k]...)
	IntakeVelocity[len(IntakeVelocity)-1] = float64(intake)
	DiffuserVelocity = append(DiffuserVelocity[k:], DiffuserVelocity[0:k]...)
	DiffuserVelocity[len(DiffuserVelocity)-1] = float64(diffuser)
	TestSectionVelocity = append(TestSectionVelocity[k:], TestSectionVelocity[0:k]...)
	TestSectionVelocity[len(TestSectionVelocity)-1] = float64(testSection)
}

func forcesToUi(f1 float32, f2 float32, f3 float32, f4 float32, f5 float32, f6 float32) {
	var k int = 1
	f1 = f1 / float32(calibration.CalibrationFactor)
	f2 = f2 / float32(calibration.CalibrationFactor)
	f3 = f3 / float32(calibration.CalibrationFactor)
	f4 = f4 / float32(calibration.CalibrationFactor)
	f5 = f5 / float32(calibration.CalibrationFactor)
	f6 = f6 / float32(calibration.CalibrationFactor)

	Strain1 = append(Strain1[k:], Strain1[0:k]...)
	Strain2 = append(Strain2[k:], Strain2[0:k]...)
	Strain3 = append(Strain3[k:], Strain3[0:k]...)
	Strain4 = append(Strain4[k:], Strain4[0:k]...)
	Strain5 = append(Strain5[k:], Strain5[0:k]...)
	Strain6 = append(Strain6[k:], Strain6[0:k]...)
	Strain1[len(Strain1)-1] = float64(f1)
	Strain2[len(Strain2)-1] = float64(f2)
	Strain3[len(Strain3)-1] = float64(f3)
	Strain4[len(Strain4)-1] = float64(f4)
	Strain5[len(Strain5)-1] = float64(f5)
	Strain6[len(Strain6)-1] = float64(f6)

	//to newtons
	f1 = f1 * 0.00981
	f2 = f2 * 0.00981
	f3 = f3 * 0.00981
	f4 = f4 * 0.00981
	f5 = f5 * 0.00981
	f6 = f6 * 0.00981
	Platform.GetForceTransformationMatrix()
	FandM := Platform.GetForceAndMoments(float64(f1), float64(f2), float64(f3), float64(f4), float64(f5), float64(f6))
	Fx = append(Fx[k:], Fx[0:k]...)
	Fx[len(Fx)-1] = FandM[0]
	Fy = append(Fy[k:], Fy[0:k]...)
	Fy[len(Fy)-1] = FandM[1]
	Fz = append(Fz[k:], Fz[0:k]...)
	Fz[len(Fz)-1] = FandM[2]
	Mx = append(Mx[k:], Mx[0:k]...)
	Mx[len(Mx)-1] = FandM[3]
	My = append(My[k:], My[0:k]...)
	My[len(My)-1] = FandM[4]
	Mz = append(Mz[k:], Mz[0:k]...)
	Mz[len(Mz)-1] = FandM[5]
	calibration.CalibrationLoads = Fz
	if RecordData && ForceMomentsWriter != nil {
		RecordToFile(ForceMomentsWriter, float32(FandM[0]), float32(FandM[1]), float32(FandM[2]), float32(FandM[3]), float32(FandM[4]), float32(FandM[5]))
	}
}

func imuToUi(yaw float32, pitch float32, roll float32) {
	var k int = 1
	Yaw = append(Yaw[k:], Yaw[0:k]...)
	Pitch = append(Pitch[k:], Pitch[0:k]...)
	Roll = append(Roll[k:], Roll[0:k]...)
	Yaw[len(Yaw)-1] = float64(yaw)
	Roll[len(Roll)-1] = float64(roll)
	Pitch[len(Pitch)-1] = float64(pitch)
}

type STDSender struct {
	conn *websocket.Conn
}

func (h *STDSender) StartOutputting(q queue.Queue) {
	fmt.Println("Start sending")
	for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {

		h.conn = api.WebsocketConn
		if h.conn == nil {
			err := fmt.Errorf("Connection not established yet")
			g.Msgbox("Error", err.Error())
			break
		}
		err := h.conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			fmt.Println("err: ", err)
			g.Msgbox("Error", err.Error())
			break
		}
		fmt.Println("Message sent: ", string(message), len(message))
	}
}

func RecordToFile(fw *fileWriter.FileWriter, vals ...float32) {
	var record []string
	record = append(record, time.Now().String())
	for val := range vals {
		record = append(record, fmt.Sprint(val))
	}
	fw.InputChannel <- record
}
