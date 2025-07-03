import {Button, MenuListItem, ScrollView, styleReset, TextInput} from "react95";
import React, {useState} from "react";
import {createGlobalStyle} from "styled-components";
import ms_sans_serif from "../assets/fonts/fixedsys.woff2";
import ms_sans_serif_bold from 'react95/dist/fonts/ms_sans_serif_bold.woff2';
import { MapContainer, TileLayer, useMap, Marker, Popup, Polyline, useMapEvents} from 'react-leaflet'
import L from 'leaflet';
import original from 'react95/dist/themes/original';

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


const zoom = 15; // Уровень масштаба
const center = [53.23292, 44.87702]
const MapClickHandler = ({ onMapClick }) => {
    useMapEvents({
        click(e) {
            onMapClick(e.latlng); // Передаем координаты клика
        },
    });
    return null;
};
const BusStopComponent = () => {
    const [markerPosition, setMarkerPosition] = useState(null);

    const handleMapClick = (latlng) => {
        setMarkerPosition(latlng); // Сохраняем координаты маркера
    };
    return (
        <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start' }}>
            <ScrollView style={{width: '200px', height: '200px', marginRight: '20px'}}>
                <MenuListItem>Днепропердовская</MenuListItem>
                <MenuListItem>Ост. Панки</MenuListItem>
                <MenuListItem>Воронины</MenuListItem>
                <MenuListItem>СМЕРТЬ БОГА</MenuListItem>

            </ScrollView>
            <div>

                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Широта</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Долгота</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Название</div>
                </div>
                <MapContainer
                    center={center} zoom={zoom}
                    scrollWheelZoom={false}
                    style={{ height: '400px', width: '100%', marginTop: '20px', marginBottom: "20px" }}
                    whenCreated={(map) => (mapRef.current = map)}
                >
                    <TileLayer
                        attribution='© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    />
                    <MapClickHandler onMapClick={handleMapClick} />
                    {markerPosition && <Marker position={markerPosition} />}
                </MapContainer>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  marginRight: "20px"}}>
                    <Button style={{marginRight: '10px'}}>Сохранить</Button>
                    <Button style={{marginRight: '10px'}}>Удалить</Button>
                    <Button >Создать</Button>
                </div>
            </div>






        </div>

    )
}

export default BusStopComponent