import React from 'react';
import {useEffect} from "react";
import {
    MenuList,
    MenuListItem,
    Separator,
    Frame,
    styleReset,
    Tabs,
    Tab,
    TabBody,
    TextInput,
    ScrollView,
    Button
} from 'react95';
import { createGlobalStyle, ThemeProvider } from 'styled-components';

import { MapContainer, TileLayer, useMap, Marker, Popup, Polyline} from 'react-leaflet'
// import 'leaflet/dist/leaflet.css'
import L from 'leaflet';
/* Pick a theme of your choice */
import original from 'react95/dist/themes/original';


/* Original Windows95 font (optional) */
import ms_sans_serif from './assets/fonts/fixedsys.woff2';
import ms_sans_serif_bold from 'react95/dist/fonts/ms_sans_serif_bold.woff2';



const GlobalStyles = createGlobalStyle`
  ${styleReset}
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif}') format('woff2');
    font-weight: 400;
    font-style: normal
  }
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif_bold}') format('woff2');
    font-weight: bold;
    font-style: normal
  }
  body, input, select, textarea {
    font-family: 'ms_sans_serif';
  }
`;

// Данные первых трех остановок автобуса №66
const stops =
    [
        {
            "Lat": "53.23292",
            "Long": "44.87702",
            "Name": "Арбековская застава"
        },
        {
            "Lat": "53.226504",
            "Long": "44.877576",
            "Name": "Бутузова"
        },
        {
            "Lat": "53.224063",
            "Long": "44.878907",
            "Name": "Школа №79"
        },
        {
            "Lat": "53.22067",
            "Long": "44.877383",
            "Name": "Стадион Запрудный"
        }
    ]


// Координаты для полилинии
const routeCoordinates = stops.map(stop => [parseFloat(stop.Lat), parseFloat(stop.Long)]);


const App = () => {
    const center = routeCoordinates[0]; // Примерно центр для первых трех остановок
    const zoom = 15; // Уровень масштаба
    return (
        <div>
            <GlobalStyles />
            <ThemeProvider theme={original}>
                <Frame
                    style={{ padding: '0.5rem', lineHeight: '1.5', width: 600 }}
                >
                    <Tabs value={"Маршруты"}>
                        <Tab value={0}>Остановки</Tab>
                        <Tab value={1}>Водители</Tab>
                        <Tab value={2}>Автобусы</Tab>
                    </Tabs>
                    <TabBody>
                        <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start' }}>
                            <ScrollView style={{width: '100px', height: '200px', marginRight: '20px'}}>
                                <MenuListItem>93</MenuListItem>
                                <MenuListItem>66</MenuListItem>
                                <MenuListItem>77</MenuListItem>
                                <MenuListItem>130</MenuListItem>

                            </ScrollView>
                            <div>
                                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                                    <div style={{marginRight: '20px'}}>Номер маршрута</div>
                                </div>
                                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                                    <ScrollView style={{width: '150px', marginRight: '20px'}}>
                                        <MenuListItem>Борщов А.Н.</MenuListItem>
                                        <MenuListItem>Иванов И. И.</MenuListItem>
                                        <MenuListItem>ХУЕСОС</MenuListItem>
                                        <MenuListItem>+</MenuListItem>
                                    </ScrollView>
                                    <div style={{marginRight: '20px'}}>Водители на маршруте</div>
                                </div>
                                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  marginBottom: '10px'}}>
                                    <ScrollView style={{width: '150px', marginRight: '20px'}}>
                                        <MenuListItem>Додеповская</MenuListItem>
                                        <MenuListItem>Хуесосовская</MenuListItem>
                                        <MenuListItem>20 лет ебания</MenuListItem>
                                        <MenuListItem>+</MenuListItem>
                                    </ScrollView>
                                    <div style={{marginRight: '20px'}}>Остановки</div>
                                </div>
                                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  marginBottom: '10px'}}>
                                    <ScrollView style={{width: '150px', marginRight: '20px'}}>
                                        <MenuListItem>Scania 666</MenuListItem>
                                        <MenuListItem>паз разъебаный</MenuListItem>
                                        <MenuListItem>Газезь</MenuListItem>
                                        <MenuListItem>+</MenuListItem>
                                    </ScrollView>
                                    <div style={{marginRight: '20px'}}>Автобусы на маршруте</div>
                                </div>
                                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  }}>
                                    <Button style={{marginRight: '10px'}}>Сохранить</Button>
                                    <Button style={{marginRight: '10px'}}>Удалить</Button>
                                    <Button >Создать</Button>
                                </div>
                            </div>
                            <MapContainer
                                center={center} zoom={zoom}
                                scrollWheelZoom={false}
                                style={{ height: '400px', width: '100%', marginTop: '20px' }}
                                whenCreated={(map) => (mapRef.current = map)}
                            >
                                <TileLayer
                                    attribution='© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                />
                                {stops.map((stop, index) => (
                                    <Marker key={index} position={[stop.Lat, stop.Long]}>
                                        <Popup>{stop.Name}</Popup>
                                    </Marker>
                                ))}

                                <Polyline positions={routeCoordinates} color="blue" weight={4} opacity={0.7} />
                            </MapContainer>
                        </div>

                    </TabBody>

                </Frame>
            </ThemeProvider>
        </div>
    )
}


export default App;